package utils

import "github.com/bwmarrin/discordgo"

type Embed struct {
	*discordgo.MessageEmbed
}

const (
	TitleLimit       = 256
	DescriptionLimit = 4096
	FieldLimit       = 25
	FieldNameLimit   = 256
	FieldValueLimit  = 1024
	FooterLimit      = 2048
	AuthorNameLimit  = 256
	Limit            = 6000
)

func NewEmbed() *Embed {
	return &Embed{&discordgo.MessageEmbed{}}
}

func (e *Embed) SetDefault() *Embed {
	e.Color = 0x3056db
	e.Footer = &discordgo.MessageEmbedFooter{
		Text:    "",
		IconURL: "",
	}
	return e
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

func (e *Embed) AddField(name string, value string, inline bool) *Embed {
	if len(value) > FieldValueLimit {
		value = value[:FieldValueLimit]
	}

	if len(name) > FieldNameLimit {
		name = name[:FieldNameLimit]
	}

	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
	return e
}

func (e *Embed) SetFooter(args ...string) *Embed {
	var (
		proxyURL string
		iconURL  string
		text     string
	)

	if len(args) == 0 {
		return e
	}

	switch {
	case len(args) > 2:
		proxyURL = args[2]
		fallthrough
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		if len(args[0]) > FooterLimit {
			text = args[0][:FooterLimit]
		} else {
			text = args[0]
		}
	}

	e.Footer = &discordgo.MessageEmbedFooter{
		IconURL:      iconURL,
		Text:         text,
		ProxyIconURL: proxyURL,
	}
	return e
}

func (e *Embed) SetImage(args ...string) *Embed {
	var (
		URL      string
		proxyURL string
	)

	if len(args) == 0 {
		return e
	}

	switch {
	case len(args) > 1:
		proxyURL = args[1]
		fallthrough
	case len(args) > 0:
		URL = args[0]
	}

	e.Image = &discordgo.MessageEmbedImage{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

func (e *Embed) SetThumbnail(args ...string) *Embed {
	var (
		URL      string
		proxyURL string
	)

	if len(args) == 0 {
		return e
	}

	switch {
	case len(args) > 1:
		proxyURL = args[1]
		fallthrough
	case len(args) > 0:
		URL = args[0]
	case len(args) == 0:
		return e
	}

	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

func (e *Embed) SetAuthor(args ...string) *Embed {
	var (
		name     string
		iconURL  string
		URL      string
		proxyURL string
	)

	if len(args) == 0 {
		return e
	}

	switch {
	case len(args) > 3:
		proxyURL = args[3]
		fallthrough
	case len(args) > 2:
		URL = args[2]
		fallthrough
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		if len(args[0]) > AuthorNameLimit {
			name = args[0][:AuthorNameLimit]
		} else {
			name = args[0]
		}
	}

	e.Author = &discordgo.MessageEmbedAuthor{
		Name:         name,
		IconURL:      iconURL,
		URL:          URL,
		ProxyIconURL: proxyURL,
	}
	return e
}

func (e *Embed) SetURL(URL string) *Embed {
	e.URL = URL
	return e
}

func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

func (e *Embed) InlineAllFields() *Embed {
	for _, field := range e.Fields {
		field.Inline = true
	}
	return e
}

func (e *Embed) UnlineAllFields() *Embed {
	for _, field := range e.Fields {
		field.Inline = false
	}
	return e
}

func (e *Embed) Truncate() *Embed {
	return e
}

func (e *Embed) TruncateFields() *Embed {
	if len(e.Fields) > FieldLimit {
		e.Fields = e.Fields[:FieldLimit]
	}

	for _, field := range e.Fields {
		if len(field.Name) > FieldNameLimit {
			field.Name = field.Name[:FieldNameLimit]
		}
		if len(field.Value) > FieldValueLimit {
			field.Value = field.Value[:FieldValueLimit]
		}
	}

	return e
}

func (e *Embed) TruncateDescription() *Embed {
	if len(e.Description) > DescriptionLimit {
		e.Description = e.Description[:DescriptionLimit]
	}
	return e
}

func (e *Embed) TruncateTitle() *Embed {
	if len(e.Title) > TitleLimit {
		e.Title = e.Title[:TitleLimit]
	}
	return e
}

func (e *Embed) TruncateFooter() *Embed {
	if e.Footer != nil && len(e.Footer.Text) > FooterLimit {
		e.Footer.Text = e.Footer.Text[:FooterLimit]
	}
	return e
}

func (e *Embed) TruncateAuthorName() *Embed {
	if e.Author != nil && len(e.Author.Name) > AuthorNameLimit {
		e.Author.Name = e.Author.Name[:AuthorNameLimit]
	}
	return e
}
