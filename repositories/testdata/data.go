package testdata

import "github.com/Naokiiiiiii/BlogApiPractice/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserID:   1,
		UserName: "naoki",
	},
	models.Article{
		ID:       2,
		Title:    "secondPost",
		Contents: "This is my second blog",
		UserID:   1,
		UserName: "naoki",
	},
	models.Article{
		ID:       3,
		Title:    "thirdPost",
		Contents: "This is my third blog",
		UserID:   1,
		UserName: "naoki",
	},
}

var ArticleInsertTestData = models.Article{
	Title:    "insertTest",
	Contents: "testest",
	UserID:   1,
}

var ArticleUpdateTestData = models.UpdateArticle{
	Title:    "Updated Title",
	Contents: "Updated Contents",
}

var CommentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st comment yeah",
	},
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "2nd comment yeah",
	},
}

var CommentInsertTestData = models.Comment{
	ArticleID: 1,
	UserID:    1,
	Message:   "CommentInsertTest",
}

var CommentUpdateTestData = models.UpdateComment{
	Message: "updateCommnet",
}

var UserTestData = models.User{
	UserID:   1,
	GoogleID: "123123123",
	UserName: "naoki",
	Email:    "exsample@gmail.com",
}

var UpdateUserTestData = models.UpdateUser{
	UserName: "updateName",
}

var NiceTestData = []models.Nice{
	models.Nice{
		NiceID:    1,
		UserID:    1,
		ArticleID: 1,
	},
}

var DeleteNiceTestData = models.Nice{
	UserID:    1,
	ArticleID: 2,
}
