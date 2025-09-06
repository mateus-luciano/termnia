package ui

import (
	"bufio"
	"io"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/mateus-luciano/termnia/internal/terminal"
)

type TerminalWidget struct {
	scroll *container.Scroll
	output *widget.Label
}

func NewTerminalWidget() *TerminalWidget {
	output := widget.NewLabel("")
	output.Wrapping = fyne.TextWrapWord
	scroll := container.NewScroll(output)
	return &TerminalWidget{
		scroll: scroll,
		output: output,
	}
}

func (tw *TerminalWidget) Append(text string) {
	tw.output.SetText(tw.output.Text + text)
}

type TerminalView struct {
	*container.Scroll
	ptmx   io.ReadWriteCloser
	widget *TerminalWidget
	writer io.WriteCloser
}

func NewTerminalView(w fyne.Window) *TerminalView {
	tw := NewTerminalWidget()
	scroll := tw.scroll

	tv := &TerminalView{Scroll: scroll, widget: tw}

	go tv.startShell(w)

	w.Canvas().SetOnTypedRune(func(r rune) {
		if tv.ptmx != nil {
			_, _ = tv.ptmx.Write([]byte(string(r)))
		}
	})

	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if tv.ptmx == nil {
			return
		}

		switch ev.Name {
		case fyne.KeyEnter:
			_, _ = tv.ptmx.Write([]byte("\r"))
		case fyne.KeyBackspace:
			_, _ = tv.ptmx.Write([]byte{0x7f})
		case fyne.KeyTab:
			_, _ = tv.ptmx.Write([]byte("\t"))
		case fyne.KeyEscape:
			_, _ = tv.ptmx.Write([]byte("\x1b"))
		case fyne.KeyUp:
			_, _ = tv.ptmx.Write([]byte("\x1b[A"))
		case fyne.KeyDown:
			_, _ = tv.ptmx.Write([]byte("\x1b[B"))
		case fyne.KeyLeft:
			_, _ = tv.ptmx.Write([]byte("\x1b[D"))
		case fyne.KeyRight:
			_, _ = tv.ptmx.Write([]byte("\x1b[C"))
		case fyne.KeyHome:
			_, _ = tv.ptmx.Write([]byte("\x1b[H"))
		case fyne.KeyEnd:
			_, _ = tv.ptmx.Write([]byte("\x1b[F"))
		case fyne.KeyPageUp:
			_, _ = tv.ptmx.Write([]byte("\x1b[5~"))
		case fyne.KeyPageDown:
			_, _ = tv.ptmx.Write([]byte("\x1b[6~"))
		case fyne.KeyDelete:
			_, _ = tv.ptmx.Write([]byte("\x1b[3~"))
		case fyne.KeyInsert:
			_, _ = tv.ptmx.Write([]byte("\x1b[2~"))
		case fyne.KeyA:
			_, _ = tv.ptmx.Write([]byte{0x01})
		case fyne.KeyC:
			_, _ = tv.ptmx.Write([]byte{0x03})
		case fyne.KeyD:
			_, _ = tv.ptmx.Write([]byte{0x04})
		case fyne.KeyZ:
			_, _ = tv.ptmx.Write([]byte{0x1A})
		}
	})

	return tv
}

func (t *TerminalView) startShell(w fyne.Window) {
	stdin, stdout, cmd, err := terminal.StartShell(120, 40)

	if err != nil {
		log.Printf("error starting shell: %v", err)

		return
	}

	t.writer = stdin

	go func() {
		r := bufio.NewReader(stdout)
		buf := make([]byte, 4096)

		for {
			n, err := r.Read(buf)
			if n > 0 {
				chunk := string(buf[:n])
				t.widget.Append(chunk)
			}

			if err != nil {
				if err == io.EOF {
					return
				}

				log.Printf("error reading shell: %v", err)

				return
			}
		}
	}()

	w.SetOnClosed(func() {
		if t.writer != nil {
			_ = t.writer.Close()
		}

		if closer, ok := stdout.(io.Closer); ok {
			_ = closer.Close()
		}

		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
	})
}
