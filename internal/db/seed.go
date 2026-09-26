package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
)

var titles = []string{
	"Building a Scalable Microservices Architecture with Go",
	"10 Essential Golang Best Practices for Clean Code",
	"Understanding Goroutines and Channels in Go",
	"How We Reduced API Latency by 40% with Redis Caching",
	"The Ultimate Guide to RESTful API Design",
	"Why We Switched from Node.js to Go for Our Backend",
	"Getting Started with Docker and Kubernetes in 2026",
	"Mastering SQL Queries: Performance Tuning & Indexing",
	"How to Implement JWT Authentication safely",
	"Top 5 Design Patterns Every Developer Should Know",
	"Exploring WebAssembly and the Future of Web Dev",
	"A Deep Dive into Database Sharding Strategies",
	"Optimizing Garbage Collection in High-Throughput Systems",
	"Building Real-Time Applications with WebSockets",
	"Why Clean Architecture Matters in Large Applications",
	"Designing Fault-Tolerant Distributed Systems",
	"Introduction to Event-Driven Architecture with Kafka",
	"Writing Effective Unit Tests in Golang",
	"How CI/CD Pipelines Speed Up Development Workflows",
	"The Pros and Cons of GraphQL vs REST",
	"Demystifying Kubernetes Controllers and Operators",
	"How to Secure Your Backend Against OWASP Top 10",
	"Benchmarking Golang HTTP Frameworks: Fiber vs Gin",
	"A Practical Guide to gRPC and Protocol Buffers",
	"Understanding System Design: Load Balancers Explained",
	"Monolith to Microservices: Lessons Learned",
	"Managing State in Modern Frontend Applications",
	"Mastering Git: Tips for Clean Commit History",
	"How to Build a Custom Monitoring Dashboard with Prometheus",
	"Navigating the World of Open Source Contribution",
	"Understanding Concurrency Models: Async/Await vs CSP",
	"How We Handled 10 Million Requests on Black Friday",
	"Effective Error Handling Patterns in Golang",
	"Designing a High-Performance Notification Service",
	"Zero-Downtime Database Migrations: A Practical Guide",
	"Demystifying TLS and HTTPS Encryption Protocols",
	"Building Cross-Platform CLI Tools with Go",
	"Why Observability is More Than Just Logging",
	"Mastering Memory Allocation and Pointers in Go",
	"An Introduction to Cloud Native Infrastructure",
	"Creating Dynamic Dashboards with React and Go",
	"Lessons Learned from Building a SaaS Platform",
	"How Domain-Driven Design (DDD) Simplifies Complex Systems",
	"Writing High-Performance Data Pipelines in Go",
	"Understanding Rate Limiting and Circuit Breakers",
	"A Beginner's Guide to Hash Tables and Hash Maps",
	"How to Profile and Benchmark Golang Applications",
	"Building Multi-Tenant SaaS Backends Efficiently",
	"The Role of AI in Modern Backend Engineering",
	"Debugging Production Outages: A Step-by-Step Guide",
}

var PostContents = []string{
	"Just shipped a new feature to production after weeks of testing! Nothing feels better than seeing those deployment pipeline checks turn green.",
	"Quick Go tip: Always handle your `http.Response.Body.Close()` immediately after checking for errors using `defer`. It prevents subtle memory leaks in long-running services.",
	"What is your go-to database choice for high-throughput read operations in 2026? We are debating between PostgreSQL with read replicas and MongoDB.",
	"Coffee status: 3 cups in. Code status: 0 bugs found, but 100 hidden bugs created. Modern software engineering at its finest.",
	"Spent the whole morning debugging a concurrent map write panic. Remember: standard Go maps are not thread-safe! Use `sync.Map` or mutexes.",
	"Finally migrated our monolithic API to microservices. Deployment times went down from 45 minutes to 3 minutes. Highly recommend the investment.",
	"Unpopular opinion: Writing clear documentation saves more developer hours than writing clever, optimized code.",
	"Working on a custom CLI tool in Golang using the Cobra library. The developer experience in Go for building system utilities is unmatched.",
	"Does anyone have good recommendations for open-source Golang monitoring dashboards? Looking for something lightweight to pair with Prometheus.",
	"Just finished reading 'Designing Data-Intensive Applications'. A must-read for any backend engineer looking to understand distributed systems.",
	"Friendly reminder to update your environment secrets and rotate your API keys regularly. Security is an ongoing process, not a one-time setup.",
	"Switched our logging library to `slog` in the standard library. Clean, structured, and zero extra dependencies needed!",
	"How do you handle background job queues in Go? Are you using Asynq, RabbitMQ, or sticking to native channels with worker pools?",
	"Just registered for the upcoming Go conference! Excited to meet fellow backend engineers and discuss system architecture.",
	"Refactoring legacy code is like playing Jenga—one wrong move and the whole production environment tumbles down.",
	"Just published a new blog post breaking down how Garbage Collection works under the hood in Golang. Link in my profile!",
	"Why is naming variables still the hardest problem in computer science? Spent 20 minutes deciding between `userID` and `usrID`.",
	"Built a simple real-time chat application over the weekend using WebSockets and Redis Pub/Sub in Go. The speed is unbelievable.",
	"When in doubt, add more unit tests. You'll thank yourself six months from now when refactoring core logic.",
	"Our team just hit 95% test coverage on our core billing service. A small win, but a huge confidence boost for deployments.",
	"Working remotely today from a local coffee shop. Sometimes a change of scenery is all you need to clear a mental block.",
	"Gopher tip: Keep your interfaces small and focused. The Go proverb says: 'The bigger the interface, the weaker the abstraction.'",
	"What are your favorite VS Code / Neovim extensions for Golang development in 2026?",
	"Successfully handled a massive traffic spike today without a single dropped request. Proper rate limiting and caching saved the day.",
	"Always benchmark before optimizing! Intuition about where code slows down is often wrong when compiled binaries are involved.",
	"Exploring WebAssembly compiled from Go for browser-side performance. Surprised at how easy the setup was.",
	"Dark mode should be mandatory on every developer platform and tool. Save the eyes!",
	"What is your strategy for graceful shutdowns in Go services? `context.WithTimeout` and `os.Interrupt` signals are essential.",
	"Just refactored an ugly `if-else` chain into a clean `switch` statement with channel select blocks. Go code readability is poetic.",
	"Automating integration tests in GitHub Actions has eliminated 90% of our post-release hotfixes.",
	"Who else loves the simplicity of `go fmt`? No arguments over tabs vs spaces—just format and move on with life.",
	"Building a multi-tenant database schema can be tricky. Do you prefer separate schemas per tenant or row-level security with a single schema?",
	"Pair programming with a senior dev today was an absolute masterclass. Always stay humble and keep learning.",
	"Saying goodbye to REST and hello to gRPC for inter-service communication. The strict protocol buffer contracts are a breath of fresh air.",
	"Pro tip for beginners: Learn to read and interpret pprof profiles early in your Go journey. It takes the guessing out of performance bottlenecks.",
	"Celebrating 1 year at my current software company today! Learnt more about backend architecture than in my entire university degree.",
	"How much time do you spend writing code versus reviewing PRs and attending architecture meetings?",
	"Never underestimate the power of a good `README.md`. A clear onboarding guide makes a repository 10x more valuable.",
	"Tried using Docker multi-stage builds today and reduced our deployment image size from 800MB to 15MB! Distroless images for the win.",
	"Late night coding session listening to synthwave. The flow state is real tonight.",
	"Docker desktop freezing right before a demo is a universal developer experience.",
	"Understanding pointer receivers versus value receivers in Go methods is key to controlling memory allocation behavior.",
	"What's your strategy for API versioning? Path prefix (`/v1/`), query params (`?version=1`), or header headers (`Accept-Version`)?",
	"Kicking off a brand new side project! Let's see if this one actually reaches production or stays in local development forever.",
	"Clean code isn't code that's short; it's code that's easy to read, modify, and reason about.",
	"Continuous integration is only as good as the reliability of your test suite. Flaky tests destroy trust in the pipeline.",
	"Spent 2 hours tracking down a bug only to realize I had a typo in an environment variable key. Standard Friday afternoon.",
	"Happy Friday everyone! May your pipelines build fast and your production environments remain stable all weekend.",
	"Zero downtime deployments using Kubernetes rolling updates feel like magic when executed properly.",
	"Always design your database schemas with future extensibility in mind, but don't over-engineer for scenarios that might never happen.",
}

var postComments = []string{
	"Great article! Really helped me understand this concept better.",
	"Thanks for sharing, this was super insightful.",
	"I completely agree with point #3!",
	"Could you elaborate more on the performance implications of this?",
	"Awesome write-up! Looking forward to part two.",
	"This saved me hours of debugging today. Thank you!",
	"Interesting perspective, though I usually take a slightly different approach.",
	"Bookmarked! Need to try this out on my current project.",
	"Nice post! What tools would you recommend for beginners?",
	"Straight to the point and very practical.",
	"How does this scale when dealing with millions of records?",
	"This is one of the clearest explanations I have read on the topic.",
	"Couldn't agree more. Clean architecture makes maintenance so much easier.",
	"Do you have a GitHub repo with an example implementation?",
	"Solid advice. I've seen these issues come up repeatedly in production.",
	"Very clear and concise. Thanks for breaking it down!",
	"Have you tried doing this with gRPC instead?",
	"Is there any significant memory overhead with this approach?",
	"This solved the exact problem I was facing this morning!",
	"Great read! Shared it with my engineering team.",
	"I'm curious how this compares to Node.js under heavy load.",
	"Really neat trick! Didn't know Golang supported this out of the box.",
	"Spot on! Adding tests early saves so much pain later.",
	"Extremely useful guide for anyone building backend services.",
	"Would love to see a follow-up post on monitoring and observability.",
	"Thanks for taking the time to write such a comprehensive breakdown.",
	"This is pure gold! Thanks for sharing your production experience.",
	"What are the main tradeoffs to consider before implementing this?",
	"Very helpful! The code examples made it super easy to follow.",
	"I had no idea this was a common anti-pattern, fixing it now!",
	"Super clean code examples. Thanks!",
	"How do you manage security and auth token refresh in this setup?",
	"Definitely sharing this with my junior devs.",
	"Brilliant explanation! Solved my confusion in 5 minutes.",
	"Do you recommend this pattern for microservices or monoliths?",
	"This is a game changer for our current architecture migration.",
	"Appreciate the real-world context instead of just theoretical examples.",
	"What database indices did you use to get those latency numbers?",
	"Simple, practical, and effective. Great post!",
	"I struggled with this concept for weeks until reading this.",
	"Thanks! Just implemented this in our staging environment.",
	"Is this compatible with the latest Golang release?",
	"Great practical tips. Keep up the good work!",
	"Would love to see a benchmark comparison on this.",
	"This is now part of our team's required reading list.",
	"Does this handle edge cases like network timeouts gracefully?",
	"Thanks for the clear step-by-step tutorial!",
	"Very relevant to what I'm working on right now.",
	"Awesome breakdown! Looking forward to your next article.",
	"Thanks for the insights, highly appreciated!",
}

var tags = []string{
	"Self Improvement", "Minimalism", "Health", "Travel", "Mindfulness",
	"Productivity", "Home Office", "Digital Detox", "Gardening", "DIY",
	"Yoga", "Sustainability", "Time Management", "Nature", "Cooking",
	"Fitness", "Personal Finance", "Writing", "Mental Health", "Learning",
}

var usernames = []string{
	"Alex Chen", "Beatrix Kiddo", "Carlos Rossi", "David Kim", "Elena Rostova",
	"Farhan Khan", "Grace Hopper", "Hannah Abbott", "Ian Wright", "Julia Santos",
	"Kavita Patel", "Liam O'Connor", "Marcus Aurelius", "Nadia Yilmaz", "Omar Hassan",
	"Priya Sharma", "Quentin Tarantino", "Rachel Green", "Samir Mehta", "Tariq Mansour",
	"Ulysses Grant", "Victoria Vance", "Will Byers", "Xavier Woods", "Yusuf Ali",
	"Zoe Saldana", "Aaron Paul", "Brianna Stewart", "Chloe Bennet", "Daniel Craig", "Daniel Craig",
	"Ethan Hunt", "Fiona Gallagher", "George Clark", "Heather Mason", "Isaac Newton",
	"Jessica Jones", "Kevin Spacey", "Laura Palmer", "Michael Scott", "Nina Williams", "Nina Williams",
	"Oscar Isaac", "Pam Beesly", "Qunn Fabray", "Ryan Gosling", "Sarah Connor",
	"Tyler Durden", "Uma Thurman", "Victor Stone", "Wanda Maximoff", "Zack Snyder",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	fmt.Printf("Seeding Started")

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

	comments := generateComments(50, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	fmt.Printf("Seeding Completed")
}

// generate posts
func generatePosts(num int, users []*store.User) []*store.Post {
	var posts []*store.Post

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		post := &store.Post{
			UserID:  user.ID,
			Content: PostContents[rand.Intn(len(PostContents))],
			Title:   titles[rand.Intn(len(titles))],
			Tags:    []string{tags[rand.Intn(len(tags))], tags[rand.Intn(len(tags))]},
		}
		posts = append(posts, post)
	}
	return posts
}

// generate users
func generateUsers(num int) []*store.User {
	var users []*store.User
	for i := 0; i < num; i++ {

		user := &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
		users = append(users, user)

	}
	return users
}

// generate comments
func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	var comments []*store.Comment

	for i := 0; i < num; i++ {
		comment := &store.Comment{
			PostId:  posts[rand.Intn(len(posts))].ID,
			UserId:  users[rand.Intn(len(users))].ID,
			Content: postComments[rand.Intn(len(postComments))],
		}
		comments = append(comments, comment)
	}
	return comments
}
