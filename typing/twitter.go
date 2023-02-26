package typing

import "time"

type TwitterAuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

type TwitterUser struct {
	ID           string    `json:"id" bson:"id"`
	Username     string    `json:"username" bson:"username"`
	AccessToken  string    `json:"access_token " bson:"access_token"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token"`
	TokenType    string    `json:"token_type" bson:"token_type"`
	RetrievedAt  time.Time `json:"retrieved_at" bson:"retrieved_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

type TwitterUserResponse struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
}

type TwitterAuthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type TwitterTweetResponse struct {
	Data struct {
		ID       string `json:"id"`
		Text     string `json:"text"`
		AuthorID string `json:"author_id"`
	} `json:"data"`
}

type TwitterTweetError struct {
	Errors []struct {
		Parameter string `json:"parameter"`
		Message   string `json:"message"`
	} `json:"errors"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Type   string `json:"type"`
}

type TwitterListTweetsResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Meta struct {
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}

type TwitterLikeTweetResponse struct {
	Data struct {
		Liked bool `json:"liked"`
	} `json:"data"`
}

type TwitterRetweetResponse struct {
	Data struct {
		Retweeted bool `json:"retweeted"`
	} `json:"data"`
}

type TwitterListResponse struct {
	Data []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		AuthorID string `json:"author_id"`
		Text     string `json:"text"`
	}
	Meta struct {
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	}
}
