package server

import (
	"go-insta-cli/pkg/posts"
)

// getMockPosts returns realistic mock Instagram posts fallback data for UI testing.
func getMockPosts() []posts.Post {
	return []posts.Post{
		{
			ID:           "17982347101",
			Caption:      "Exploring the beauty of urban architecture. Sunset vibes in the city skyline! #cityscape #photography #sunset",
			Timestamp:    "2026-10-07T18:30:00Z",
			LikeCount:    245,
			CommentCount: 18,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1513694203232-719a280e022f?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347102",
			Caption:      "Morning coffee and productive coding sessions. Building modern CLI tools with Go! ☕💻 #golang #developer #coding",
			Timestamp:    "2026-10-06T09:15:00Z",
			LikeCount:    512,
			CommentCount: 42,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1517694712202-14dd9538aa97?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347103",
			Caption:      "Weekend getaway into nature. Fresh mountain air and quiet trails. 🏔️🌲 #nature #hiking #adventure",
			Timestamp:    "2026-10-04T14:20:00Z",
			LikeCount:    890,
			CommentCount: 65,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1464822759023-fed622ff2c3b?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347104",
			Caption:      "Minimalist workspace setup for maximum focus. Clean desk, clear mind. ✨ #workspace #setup #minimalism",
			Timestamp:    "2026-10-02T11:45:00Z",
			LikeCount:    378,
			CommentCount: 29,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1499750310107-5fef28a66643?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347105",
			Caption:      "Quick preview of our upcoming project UI dashboard! Stay tuned for more updates. 🚀",
			Timestamp:    "2026-09-30T16:00:00Z",
			LikeCount:    1024,
			CommentCount: 94,
			MediaType:    "VIDEO",
			MediaURL:     "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347106",
			Caption:      "Delicious homemade ramen for dinner. Perfect dish for a cozy evening. 🍜 #foodie #ramen #cooking",
			Timestamp:    "2026-09-28T20:10:00Z",
			LikeCount:    630,
			CommentCount: 37,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1569718212165-3a8278d5f624?w=600&auto=format&fit=crop&q=80",
		},
	}
}
