package discord

import "github.com/Goscord/goscord/discord/embed"

type ApplicationCommandType int

const (
	_                      ApplicationCommandType = iota
	ApplicationCommandChat                        // slash command
	ApplicationCommandUser
	ApplicationCommandMessage
)

type ApplicationCommandOptionType int

const (
	_ ApplicationCommandOptionType = iota
	ApplicationCommandOptionSubCommand
	ApplicationCommandOptionSubCommandGroup
	ApplicationCommandOptionString
	ApplicationCommandOptionInteger
	ApplicationCommandOptionBoolean
	ApplicationCommandOptionUser
	ApplicationCommandOptionChannel
	ApplicationCommandOptionRole
	ApplicationCommandOptionMentionable
	ApplicationCommandOptionNumber
	ApplicationCommandOptionAttachment
)

type ApplicationCommandPermissionType int

const (
	_ ApplicationCommandPermissionType = iota
	ApplicationCommandPermissionTypeRole
	ApplicationCommandPermissionTypeUser
	ApplicationCommandPermissionTypeChannel
)

type ComponentType int

const (
	ComponentTypeActionRow  ComponentType = iota + 1 // container for other components
	ComponentTypeButton                              // clickable button
	ComponentTypeSelectMenu                          // dropdown
	ComponentTypeTextInput                           // text input
)

type ApplicationCommand struct {
	Id                       string                 `json:"id,omitempty"`
	ApplicationId            string                 `json:"application_id,omitempty"`
	GuildId                  string                 `json:"guild_id,omitempty"`
	Version                  string                 `json:"version,omitempty"`
	Type                     ApplicationCommandType `json:"type,omitempty"`
	Name                     string                 `json:"name"`
	NameLocalizations        map[Locale]string      `json:"name_localizations,omitempty"`
	DefaultPermission        bool                   `json:"default_permission,omitempty"`
	DefaultMemberPermissions int64                  `json:"default_member_permissions,string,omitempty"`
	DMPermission             bool                   `json:"dm_permission,omitempty"`

	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Locale]string           `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options"`
}

type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[Locale]string `json:"name_localizations,omitempty"`
	Value             interface{}       `json:"value"`
}

type ApplicationCommandPermissions struct {
	Id         string                           `json:"id"`
	Type       ApplicationCommandPermissionType `json:"type"`
	Permission bool                             `json:"permission"`
}

type ApplicationCommandPermissionsList struct {
	Permissions []*ApplicationCommandPermissions `json:"permissions"`
}

type GuildApplicationCommandPermissions struct {
	Id            string                           `json:"id"`
	ApplicationId string                           `json:"application_id"`
	GuilddId      string                           `json:"guild_id"`
	Permissions   []*ApplicationCommandPermissions `json:"permissions"`
}

type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Type    ApplicationCommandOptionType               `json:"type"`
	Value   interface{}                                `json:"value,omitempty"` // string, integer, or double
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused bool                                       `json:"focused,omitempty"`
}

func (o ApplicationCommandInteractionDataOption) Int() int64 {
	if o.Type != ApplicationCommandOptionInteger {
		return 0
	}

	return int64(o.Value.(float64))
}

func (o ApplicationCommandInteractionDataOption) Float() float64 {
	if o.Type != ApplicationCommandOptionNumber {
		return 0
	}

	return o.Value.(float64)
}

func (o ApplicationCommandInteractionDataOption) String() string {
	if o.Type != ApplicationCommandOptionString {
		return ""
	}

	return o.Value.(string)
}

func (o ApplicationCommandInteractionDataOption) Bool() bool {
	if o.Type != ApplicationCommandOptionBoolean {
		return false
	}

	return o.Value.(bool)
}

type ApplicationCommandOption struct {
	Type                     ApplicationCommandOptionType      `json:"type"`
	Name                     string                            `json:"name"`
	NameLocalizations        map[Locale]string                 `json:"name_localizations,omitempty"`
	Description              string                            `json:"description,omitempty"`
	DescriptionLocalizations map[Locale]string                 `json:"description_localizations,omitempty"`
	ChannelTypes             []int                             `json:"channel_types"`
	Required                 bool                              `json:"required"`
	Options                  []*ApplicationCommandOption       `json:"options"`
	Autocomplete             bool                              `json:"autocomplete"`
	Choices                  []*ApplicationCommandOptionChoice `json:"choices"`
	MinValue                 float64                           `json:"min_value,omitempty"`
	MaxValue                 float64                           `json:"max_value,omitempty"`
	MinLength                int                               `json:"min_length,omitempty"`
	MaxLength                int                               `json:"max_length,omitempty"`
}

// Interaction
type InteractionType int

const (
	_ InteractionType = iota
	InteractionTypePing
	InteractionTypeApplicationCommand
	InteractionTypeMessageComponent
	InteractionTypeApplicationCommandAutocomplete
	InteractionTypeModalSubmit
)

type InteractionCallbackType int

const (
	InteractionCallbackTypePong                                 = 1 // ack a ping
	InteractionCallbackTypeChannelWithSource                    = 4 // respond to an interaction with a message
	InteractionCallbackTypeDeferredChannelMessageWithSource     = 5 // ACK an interaction and edit a response later, the user sees a loading state
	InteractionCallbackTypeDeferredUpdateMessage                = 6 // for components, ACK an interaction and edit the original message later; the user does not see a loading state
	InteractionCallbackTypeUpdateMessage                        = 7 // for components, edit the message the component was attached to
	InteractionCallbackTypeApplicationCommandAutocompleteResult = 8 // for autocomplete, return the results of the autocomplete
	InteractionCallbackTypeModal                                = 9 // respond to an interaction with a popup modal
)

type Interaction struct {
	Id             string                `json:"id"`
	ApplicationId  string                `json:"application_id"`
	Type           InteractionType       `json:"type"`
	Data           *InteractionData      `json:"data,omitempty"`
	GuildId        string                `json:"guild_id,omitempty"`
	ChannelId      string                `json:"channel_id,omitempty"`
	Member         *GuildMember          `json:"member,omitempty"`
	User           *User                 `json:"user,omitempty"`
	Token          string                `json:"token"`
	Version        int                   `json:"version"`
	Message        *Message              `json:"message,omitempty"`
	AppPermissions BitwisePermissionFlag `json:"app_permissions,string,omitempty"`
	Locale         Locale                `json:"locale,omitempty"`
	GuildLocale    Locale                `json:"guild_locale,omitempty"`
}

type InteractionData struct {
	Id       string                                     `json:"id"`
	Name     string                                     `json:"name"`
	Type     ApplicationCommandType                     `json:"type"`
	Resolved *ResolvedData                              `json:"resolved,omitempty"`
	Options  []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	GuildId  string                                     `json:"guild_id,omitempty"`
	TargetId string                                     `json:"target_id,omitempty"` // user or message
}

type MessageComponent struct {
	CustomId      string        `json:"custom_id"`
	ComponentType ComponentType `json:"component_type"`
	//Values map[string]interface{} `json:"values,omitempty"`
}

type ModalSubmit struct {
	CustomId   string              `json:"custom_id"`
	Components []*MessageComponent `json:"components"`
}

type ResolvedData struct {
	Users       []*User        `json:"users,omitempty"`
	Members     []*GuildMember `json:"members,omitempty"`
	Roles       []*Role        `json:"roles,omitempty"`
	Channels    []*Channel     `json:"channels,omitempty"`
	Messages    []*Message     `json:"messages,omitempty"`
	Attachments []*Attachment  `json:"attachments,omitempty"`
}

type MessageInteraction struct {
	Id     string          `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   *User           `json:"user"`
	Member *GuildMember    `json:"member,omitempty"`
}

type InteractionResponse struct {
	Type InteractionCallbackType `json:"type"`
	Data interface{}             `json:"data"` // depends on type
}

type InteractionCallbackMessage struct {
	Tts             bool                `json:"tts,omitempty"`
	Content         string              `json:"content,omitempty"`
	Embeds          []*embed.Embed      `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions    `json:"allowed_mentions,omitempty"`
	Flags           MessageFlag         `json:"flags,omitempty"`
	Components      []*MessageComponent `json:"components,omitempty"`
	Attachments     []*Attachment       `json:"attachments,omitempty"`
}

type InteractionCallbackAutocomplete struct {
	Choices []*ApplicationCommandOptionChoice `json:"choices"`
}

type InteractionCallbackModal struct {
	CustomId   string              `json:"custom_id"` // a developer-defined identifier for the component, max 100 characters
	Title      string              `json:"title"`
	Components []*MessageComponent `json:"components"`
}
