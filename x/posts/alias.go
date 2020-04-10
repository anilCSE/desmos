package posts

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/common"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/polls"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/reactions"
	"github.com/desmos-labs/desmos/x/posts/internal/types/msgs"
)

const (
	ModuleName                      = common.ModuleName
	RouterKey                       = common.RouterKey
	StoreKey                        = common.StoreKey
	MaxPostMessageLength            = common.MaxPostMessageLength
	MaxOptionalDataFieldsNumber     = common.MaxOptionalDataFieldsNumber
	MaxOptionalDataFieldValueLength = common.MaxOptionalDataFieldValueLength
	ActionCreatePost                = common.ActionCreatePost
	ActionEditPost                  = common.ActionEditPost
	ActionAnswerPoll                = common.ActionAnswerPoll
	ActionAddPostReaction           = common.ActionAddPostReaction
	ActionRemovePostReaction        = common.ActionRemovePostReaction
	ActionRegisterReaction          = common.ActionRegisterReaction
	QuerierRoute                    = common.QuerierRoute
	QueryPost                       = common.QueryPost
	QueryPosts                      = common.QueryPosts
	QueryPollAnswers                = common.QueryPollAnswers
	QueryRegisteredReactions        = common.QueryRegisteredReactions
	PostSortByID                    = common.PostSortByID
	PostSortByCreationDate          = common.PostSortByCreationDate
	PostSortOrderAscending          = common.PostSortOrderAscending
	PostSortOrderDescending         = common.PostSortOrderDescending
	OpWeightMsgCreatePost           = simulation.OpWeightMsgCreatePost
	OpWeightMsgEditPost             = simulation.OpWeightMsgEditPost
	OpWeightMsgAddReaction          = simulation.OpWeightMsgAddReaction
	OpWeightMsgRemoveReaction       = simulation.OpWeightMsgRemoveReaction
	OpWeightMsgAnswerPoll           = simulation.OpWeightMsgAnswerPoll
	OpWeightMsgRegisterReaction     = simulation.OpWeightMsgRegisterReaction
	DefaultGasValue                 = simulation.DefaultGasValue
	EventTypePostCreated            = types.EventTypePostCreated
	EventTypePostEdited             = types.EventTypePostEdited
	EventTypePostReactionAdded      = types.EventTypePostReactionAdded
	EventTypePostReactionRemoved    = types.EventTypePostReactionRemoved
	EventTypeAnsweredPoll           = types.EventTypeAnsweredPoll
	EventTypeClosePoll              = types.EventTypeClosePoll
	EventTypeRegisterReaction       = types.EventTypeRegisterReaction
	AttributeKeyPostID              = types.AttributeKeyPostID
	AttributeKeyPostParentID        = types.AttributeKeyPostParentID
	AttributeKeyPostOwner           = types.AttributeKeyPostOwner
	AttributeKeyPostEditTime        = types.AttributeKeyPostEditTime
	AttributeKeyPollAnswerer        = types.AttributeKeyPollAnswerer
	AttributeKeyPostReactionOwner   = types.AttributeKeyPostReactionOwner
	AttributeKeyPostReactionValue   = types.AttributeKeyPostReactionValue
	AttributeKeyReactionCreator     = types.AttributeKeyReactionCreator
	AttributeKeyReactionShortCode   = types.AttributeKeyReactionShortCode
	AttributeKeyReactionSubSpace    = types.AttributeKeyReactionSubSpace
	AttributeKeyCreationTime        = types.AttributeKeyCreationTime
)

var (
	// functions aliases
	RegisterCodec                 = types.RegisterCodec
	NewGenesisState               = types.NewGenesisState
	DefaultGenesisState           = types.DefaultGenesisState
	ValidateGenesis               = types.ValidateGenesis
	DefaultQueryPostsParams       = types.DefaultQueryPostsParams
	ParsePostID                   = models.ParsePostID
	NewPostResponse               = models.NewPostResponse
	RegisterModelsCodec           = models.RegisterModelsCodec
	NewPostCreationData           = models.NewPostCreationData
	PostStoreKey                  = models.PostStoreKey
	PostCommentsStoreKey          = models.PostCommentsStoreKey
	PostReactionsStoreKey         = models.PostReactionsStoreKey
	ReactionsStoreKey             = models.ReactionsStoreKey
	PollAnswersStoreKey           = models.PollAnswersStoreKey
	NewPost                       = models.NewPost
	NewPostMedia                  = common.NewPostMedia
	ValidateURI                   = common.ValidateURI
	NewPostMedias                 = common.NewPostMedias
	ParseAnswerID                 = polls.ParseAnswerID
	NewPollAnswer                 = polls.NewPollAnswer
	NewPollAnswers                = polls.NewPollAnswers
	NewPollData                   = polls.NewPollData
	ArePollDataEquals             = polls.ArePollDataEquals
	NewUserAnswer                 = polls.NewUserAnswer
	NewPostReaction               = reactions.NewPostReaction
	NewReaction                   = reactions.NewReaction
	IsEmoji                       = reactions.IsEmoji
	NewMsgRegisterReaction        = msgs.NewMsgRegisterReaction
	RegisterMessagesCodec         = msgs.RegisterMessagesCodec
	NewMsgCreatePost              = msgs.NewMsgCreatePost
	NewMsgEditPost                = msgs.NewMsgEditPost
	NewMsgAnswerPoll              = msgs.NewMsgAnswerPoll
	NewMsgAddPostReaction         = msgs.NewMsgAddPostReaction
	NewMsgRemovePostReaction      = msgs.NewMsgRemovePostReaction
	NewKeeper                     = keeper.NewKeeper
	NewQuerier                    = keeper.NewQuerier
	NewHandler                    = keeper.NewHandler
	HandlePostCreationRequest     = keeper.HandlePostCreationRequest
	SimulateMsgAnswerToPoll       = simulation.SimulateMsgAnswerToPoll
	SimulateMsgCreatePost         = simulation.SimulateMsgCreatePost
	SimulateMsgEditPost           = simulation.SimulateMsgEditPost
	SimulateMsgAddPostReaction    = simulation.SimulateMsgAddPostReaction
	SimulateMsgRemovePostReaction = simulation.SimulateMsgRemovePostReaction
	SimulateMsgRegisterReaction   = simulation.SimulateMsgRegisterReaction
	RandomPost                    = simulation.RandomPost
	RandomPostData                = simulation.RandomPostData
	RandomPostReactionData        = simulation.RandomPostReactionData
	RandomPostReactionValue       = simulation.RandomPostReactionValue
	RandomPostID                  = simulation.RandomPostID
	RandomMessage                 = simulation.RandomMessage
	RandomSubspace                = simulation.RandomSubspace
	RandomHashtag                 = simulation.RandomHashtag
	RandomMedias                  = simulation.RandomMedias
	RandomPollData                = simulation.RandomPollData
	GetAccount                    = simulation.GetAccount
	RandomReactionValue           = simulation.RandomReactionValue
	RandomReactionShortCode       = simulation.RandomReactionShortCode
	RandomReactionData            = simulation.RandomReactionData
	RegisteredReactionsData       = simulation.RegisteredReactionsData
	DecodeStore                   = simulation.DecodeStore
	RandomizedGenState            = simulation.RandomizedGenState
	WeightedOperations            = simulation.WeightedOperations

	// variable aliases
	MsgsCodec                = msgs.MsgsCodec
	RandomMimeTypes          = simulation.RandomMimeTypes
	RandomHosts              = simulation.RandomHosts
	ModuleCdc                = types.ModuleCdc
	ModelsCdc                = models.ModelsCdc
	SubspaceRegEx            = common.SubspaceRegEx
	HashtagRegEx             = common.HashtagRegEx
	ShortCodeRegEx           = common.ShortCodeRegEx
	URIRegEx                 = common.URIRegEx
	LastPostIDStoreKey       = common.LastPostIDStoreKey
	PostStorePrefix          = common.PostStorePrefix
	PostCommentsStorePrefix  = common.PostCommentsStorePrefix
	PostReactionsStorePrefix = common.PostReactionsStorePrefix
	ReactionsStorePrefix     = common.ReactionsStorePrefix
	PollAnswersStorePrefix   = common.PollAnswersStorePrefix
)

type (
	PostID                   = models.PostID
	PostIDs                  = models.PostIDs
	PostQueryResponse        = models.PostQueryResponse
	PostCreationData         = models.PostCreationData
	PollAnswersQueryResponse = models.PollAnswersQueryResponse
	Post                     = models.Post
	Posts                    = models.Posts
	PostMedia                = common.PostMedia
	PostMedias               = common.PostMedias
	OptionalData             = common.OptionalData
	KeyValue                 = common.KeyValue
	AnswerID                 = polls.AnswerID
	PollAnswer               = polls.PollAnswer
	PollAnswers              = polls.PollAnswers
	PollData                 = polls.PollData
	UserAnswer               = polls.UserAnswer
	UserAnswers              = polls.UserAnswers
	PostReaction             = reactions.PostReaction
	PostReactions            = reactions.PostReactions
	Reaction                 = reactions.Reaction
	Reactions                = reactions.Reactions
	MsgRegisterReaction      = msgs.MsgRegisterReaction
	MsgCreatePost            = msgs.MsgCreatePost
	MsgEditPost              = msgs.MsgEditPost
	MsgAnswerPoll            = msgs.MsgAnswerPoll
	MsgAddPostReaction       = msgs.MsgAddPostReaction
	MsgRemovePostReaction    = msgs.MsgRemovePostReaction
	Keeper                   = keeper.Keeper
	PostData                 = simulation.PostData
	PostReactionData         = simulation.PostReactionData
	ReactionData             = simulation.ReactionData
	GenesisState             = types.GenesisState
	QueryPostsParams         = types.QueryPostsParams
)
