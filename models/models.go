package models

import (
	"time"

	"github.com/google/uuid"
	database "github.com/niltongc/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(dbFeed))
	}

	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func DatabaseFeedFollowToFeedFollow(dbFeedFollow database.FeedsFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedsFollow) []FeedFollow {
	feedfollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedfollows = append(feedfollows, DatabaseFeedFollowToFeedFollow(dbFeedFollow))
	}

	return feedfollows
}

type Post struct {
	ID          uuid.UUID     `json:"id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Title       string        `json:"title"`
	Url         string        `json:"url"`
	Description *string       `json:"description"`
	PublishedAt time.Time     `json:"published_at"`
	FeedID      uuid.NullUUID `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string

	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		FeedID:      dbPost.FeedID,
	}
}

func DatabasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}

// func NullTimeToTimePtr(t sql.NullTime) *time.Time {
// 	if t.Valid {
// 		return &t.Time
// 	}
// 	return nil
// }

// func NullStringToStringPtr(s sql.NullString) *string {
// 	if s.Valid {
// 		return &s.String
// 	}
// 	return nil
// }
