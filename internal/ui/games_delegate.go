package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type gameItemDelegate struct{}

func (d gameItemDelegate) Height() int                             { return 1 }
func (d gameItemDelegate) Spacing() int                            { return 0 }
func (d gameItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d gameItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i := listItem.(gameInfoItem).value
	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}
	_, err := fmt.Fprint(w, fn(str))
	if err != nil {
		return
	}
}
