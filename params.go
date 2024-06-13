package getstream

type DeleteChannelParams struct {
	HardDelete *bool `json:"hard_delete,omitempty"`
}

type DeleteFileParams struct {
	Url *string `json:"url,omitempty"`
}

type DeleteImageParams struct {
	Url *string `json:"url,omitempty"`
}

type GetManyMessagesParams struct {
	Ids []string `json:"ids"`
}

type QueryMembersParams struct {
	Payload *QueryMembersRequest `json:"payload,omitempty"`
}

type DeleteMessageParams struct {
	DeletedBy *string `json:"deleted_by,omitempty"`

	Hard *bool `json:"hard,omitempty"`
}

type GetMessageParams struct {
	ShowDeletedMessage *bool `json:"show_deleted_message,omitempty"`
}

type DeleteReactionParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type GetReactionsParams struct {
	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`
}

type RemovePollVoteParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type GetRepliesParams struct {
	CreatedAtAfter *Timestamp `json:"created_at_after,omitempty"`

	CreatedAtAfterOrEqual *Timestamp `json:"created_at_after_or_equal,omitempty"`

	CreatedAtAround *Timestamp `json:"created_at_around,omitempty"`

	CreatedAtBefore *Timestamp `json:"created_at_before,omitempty"`

	CreatedAtBeforeOrEqual *Timestamp `json:"created_at_before_or_equal,omitempty"`

	IdAround *string `json:"id_around,omitempty"`

	IdGt *string `json:"id_gt,omitempty"`

	IdGte *string `json:"id_gte,omitempty"`

	IdLt *string `json:"id_lt,omitempty"`

	IdLte *string `json:"id_lte,omitempty"`

	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`
}

type QueryMessageFlagsParams struct {
	Payload *QueryMessageFlagsRequest `json:"payload,omitempty"`
}

type QueryPollsParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type DeletePollParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type GetPollParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type DeletePollOptionParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type GetPollOptionParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type QueryPollVotesParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type QueryBannedUsersParams struct {
	Payload *QueryBannedUsersRequest `json:"payload,omitempty"`
}

type SearchParams struct {
	Payload *SearchRequest `json:"payload,omitempty"`
}

type GetThreadParams struct {
	ConnectionId *string `json:"connection_id,omitempty"`

	MemberLimit *int `json:"member_limit,omitempty"`

	ParticipantLimit *int `json:"participant_limit,omitempty"`

	ReplyLimit *int `json:"reply_limit,omitempty"`
}

type DeleteDeviceParams struct {
	Id string `json:"id"`

	UserId *string `json:"user_id,omitempty"`
}

type ListDevicesParams struct {
	UserId *string `json:"user_id,omitempty"`
}

type UnbanParams struct {
	TargetUserId string `json:"target_user_id"`

	ChannelCid *string `json:"channel_cid,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`
}

type GetOGParams struct {
	Url string `json:"url"`
}

type GetRateLimitsParams struct {
	Android *bool `json:"android,omitempty"`

	Endpoints *string `json:"endpoints,omitempty"`

	Ios *bool `json:"ios,omitempty"`

	ServerSide *bool `json:"server_side,omitempty"`

	Web *bool `json:"web,omitempty"`
}

type QueryUsersParams struct {
	Payload *QueryUsersPayload `json:"payload,omitempty"`
}

type GetCallParams struct {
	MembersLimit *int `json:"members_limit,omitempty"`

	Notify *bool `json:"notify,omitempty"`

	Ring *bool `json:"ring,omitempty"`
}