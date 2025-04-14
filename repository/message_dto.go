package repository

import "message/types/message"

type PaginatedMessageSearch struct {
    Page    int `json:"page"`
    Limit   int `json:"limit"`
    Total   int `json:"total"`
    Results []message.Message `json:"results"`
}
