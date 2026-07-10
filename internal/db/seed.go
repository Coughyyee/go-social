package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/coughyyee/social/internal/store"
)

var usernames = []string{
	"Olivia",
	"Liam",
	"Emma",
	"Noah",
	"Amelia",
	"Oliver",
	"Sophia",
	"Elijah",
	"Isla",
	"James",
	"Ava",
	"Lucas",
	"Mia",
	"Henry",
	"Grace",
	"Leo",
	"Freya",
	"Jack",
	"Lily",
	"Charlie",
	"Ella",
	"Theo",
	"Evie",
	"Oscar",
	"Ruby",
	"Arthur",
	"Chloe",
	"Thomas",
	"Ivy",
	"William",
	"Hannah",
	"George",
	"Rosie",
	"Harry",
	"Sienna",
	"Finn",
	"Alice",
	"Jacob",
	"Matilda",
	"Archie",
	"Florence",
	"Ethan",
	"Poppy",
	"Isaac",
	"Maisie",
	"Daniel",
	"Willow",
	"Samuel",
	"Eleanor",
	"Benjamin",
}

var titles = []string{
	"Learning Go",
	"My First API",
	"Morning Coffee",
	"Weekend Plans",
	"Debugging Tips",
	"Healthy Habits",
	"Travel Dreams",
	"Book Review",
	"Daily Workout",
	"New Project",
	"Simple Recipes",
	"Tech News",
	"Code Refactoring",
	"Productivity Hacks",
	"Weekend Hike",
	"Music Playlist",
	"Career Goals",
	"Learning Docker",
	"Quick Thoughts",
	"Building an App",
}

var contents = []string{
	"Today I learned something new and wanted to share it with everyone.",
	"Small improvements every day lead to big results over time.",
	"Just finished working on a fun side project this afternoon.",
	"Go continues to impress me with its simplicity and performance.",
	"Sometimes taking a short break is the best way to solve a problem.",
	"Reading documentation carefully often saves hours of debugging.",
	"Coffee, code, and good music make for a productive morning.",
	"Spent some time cleaning up old code and it feels much better now.",
	"Excited to keep learning and building new things this week.",
	"Finally fixed a bug that had been bothering me all day.",
	"Trying out a different workflow has made development much smoother.",
	"Consistency is far more valuable than occasional bursts of motivation.",
	"I've been experimenting with new tools to improve my productivity.",
	"Every project teaches something useful, even when things go wrong.",
	"Building software is all about solving problems one step at a time.",
	"Today's goal was simple: write cleaner, more maintainable code.",
	"I've started documenting my projects more thoroughly and it's paying off.",
	"Nothing beats the feeling of seeing everything compile successfully.",
	"Learning by building real projects is still my favorite approach.",
	"Looking forward to what I'll create next.",
}

var tags = []string{
	"go",
	"golang",
	"programming",
	"coding",
	"software",
	"developer",
	"webdev",
	"backend",
	"frontend",
	"api",
	"database",
	"postgres",
	"mysql",
	"docker",
	"kubernetes",
	"cloud",
	"aws",
	"linux",
	"opensource",
	"github",
	"javascript",
	"typescript",
	"react",
	"nextjs",
	"nodejs",
	"python",
	"rust",
	"cpp",
	"java",
	"dotnet",
	"testing",
	"devops",
	"microservices",
	"concurrency",
	"security",
	"performance",
	"algorithms",
	"datastructures",
	"machinelearning",
	"ai",
	"productivity",
	"tutorial",
	"beginners",
	"learning",
	"career",
	"tech",
	"startup",
	"design",
	"architecture",
	"opensource",
}

var comments = []string{
	"Great post!",
	"Really helpful.",
	"Nice work!",
	"Thanks for sharing.",
	"Love this!",
	"Very informative.",
	"Well explained.",
	"Awesome project!",
	"Keep it up!",
	"This is cool.",
	"I learned something new.",
	"Completely agree.",
	"Interesting perspective.",
	"Looking forward to more!",
	"Great advice.",
	"This was useful.",
	"Nicely written.",
	"Fantastic read.",
	"Thanks for the insight.",
	"Excellent content!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding completed")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Version: 0,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
