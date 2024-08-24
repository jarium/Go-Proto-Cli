package content

import "fmt"

type Builder struct {
	content string
}

func NewBuilder() *Builder {
	return &Builder{}
}

// Add entry to content
func (b *Builder) Add(s string) {
	b.content += s
}

// Addln adds entry to content with new line tag on suffix
func (b *Builder) Addln(s string) {
	b.content += s + "\n"
}

// Addf adds entry via format to content
func (b *Builder) Addf(format string, a ...any) {
	b.content += fmt.Sprintf(format, a...)
}

// Addfln adds entry via format, with a new line tag on suffix to content
func (b *Builder) Addfln(format string, a ...any) {
	b.content += fmt.Sprintf(format, a...) + "\n"
}

// Br adds new lines to content equal to count
func (b *Builder) Br(count int) {
	for i := 0; i < count; i++ {
		b.content += "\n"
	}
}

// Get builded content
func (b *Builder) Get() string {
	return b.content
}
