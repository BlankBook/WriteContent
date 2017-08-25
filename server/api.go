package server

import (
    "time"
    "strconv"
    "net/http"
    "database/sql"

    "github.com/blankbook/shared/models"
    "github.com/blankbook/shared/web"
)

// SetupAPI adds the API routes to the provided router
func SetupAPI(r web.Router, db *sql.DB) {
    r.HandleRoute([]string{web.POST}, "/post",
                  []string{}, []string{},
                  PostPost, db)
    // NEXT: Add 'PUT /post' to modify post content
    r.HandleRoute([]string{web.PUT}, "/post/vote",
                  []string{"userid", "postid", "vote"}, []string{},
                  PutPostVote, db)
    r.HandleRoute([]string{web.POST}, "/post/comment",
                  []string{}, []string{},
                  PostPostComment, db)
    r.HandleRoute([]string{web.PUT}, "/post/comment/vote",
                  []string{"userid", "postid", "commentid", "vote"}, []string{},
                  PutPostCommentVote, db)
    r.HandleRoute([]string{web.GET}, "/health",
                  []string{}, []string{},
                  GetHealth, db) 
}

func PostPost(w http.ResponseWriter, q map[string][]string, b string, db *sql.DB) {
    var err error
    defer func() {
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
        }
    }()
    p, err := models.ParsePost(b)
    if err != nil {
        return
    }
    if p.Time == 0 {
        p.Time = time.Now().Unix()
    }
    p.Rank = -1
    err = p.Validate()
    if err != nil {
        return
    }
    query :=`
        INSERT INTO Posts 
        (Title, Content, ContentType, GroupName, Time, Color)
        Values ($1, $2, $3, $4, $5, $6)`

    _, err = db.Exec(query, p.Title, p.Content, p.ContentType, p.GroupName, p.Time, p.Color)
    if err != nil {
        return
    }
    w.WriteHeader(http.StatusOK)
}

func PutPostVote(w http.ResponseWriter, q map[string][]string, b string, db *sql.DB) {
    userId := q["userid"][0]
    postId := q["postid"][0]
    vote, err := strconv.Atoi(q["vote"][0])
    if vote > 1 || vote < -1 || err != nil {
        http.Error(w, "Vote must be from -1 to 1", http.StatusBadRequest)
        return
    }
    query := `
        DECLARE @Values TABLE (Value INT);
        SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
        BEGIN TRANSACTION;
        UPDATE TOP (1) dbo.postVotes SET Value=$3
        OUTPUT DELETED.Value INTO @Values
        WHERE VoterID=$1 AND PostID=$2;
        IF @@ROWCOUNT = 0
            BEGIN
            INSERT dbo.postVotes (VoterID, PostID, Value) SELECT $1, $2, $3;
            UPDATE TOP (1) dbo.posts SET Score=Score+$3 WHERE ID=$2
            END
        ELSE
            BEGIN
            DECLARE @Value INT 
            SELECT TOP 1 @Value=Value
            FROM @Values
            UPDATE TOP (1) dbo.posts SET Score=Score+$3-@Value WHERE ID=$2
            END
        COMMIT TRANSACTION;`

    _, err = db.Query(query, userId, postId, vote)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return;
    }
    w.WriteHeader(http.StatusOK)
}

func PostPostComment(w http.ResponseWriter, q map[string][]string, b string, db *sql.DB) {
    var err error
    defer func() {
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
        }
    }()
    c, err := models.ParseComment(b)
    if err != nil {
        return
    }
    err = c.Validate()
    if err != nil {
        return
    }
    query :=`
        INSERT INTO Comments
        (Score, ParentPost, ParentComment, Content, EditContent, Time, Color)
        Values ($1, $2, $3, $4, $5, $6, $7)`

    _, err = db.Exec(query, c.Score, c.ParentPost, c.ParentComment, c.Content,
                     c.EditContent, c.Time, c.Color)
    if err != nil {
        return
    }
    w.WriteHeader(http.StatusOK)
}

func PutPostCommentVote(w http.ResponseWriter, q map[string][]string, b string, db *sql.DB) {
    userId := q["userid"][0]
    postId := q["postid"][0]
    commentId := q["commentid"][0]
    vote, err := strconv.Atoi(q["vote"][0])
    if vote > 1 || vote < -1 || err != nil {
        http.Error(w, "Vote must be from -1 to 1", http.StatusBadRequest)
        return
    }
    query := `
        DECLARE @Values TABLE (Value INT);
        SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
        BEGIN TRANSACTION;
        UPDATE TOP (1) dbo.commentVotes SET Value=$4
        OUTPUT DELETED.Value INTO @Values
        WHERE VoterID=$1 AND PostID=$2 AND CommentID=$3;
        IF @@ROWCOUNT = 0
            BEGIN
            INSERT dbo.commentVotes (VoterID, PostID, CommentID, Value) SELECT $1, $2, $3, $4;
            UPDATE TOP (1) dbo.comments SET Score=Score+$4 WHERE ID=$3
            END
        ELSE
            BEGIN
            DECLARE @Value INT 
            SELECT TOP 1 @Value=Value
            FROM @Values
            UPDATE TOP (1) dbo.comments SET Score=Score+$4-@Value WHERE ID=$3
            END
        COMMIT TRANSACTION;`

    _, err = db.Query(query, userId, postId, commentId, vote)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return;
    }
    w.WriteHeader(http.StatusOK)
}

func GetHealth(w http.ResponseWriter, q map[string][]string, b string, db *sql.DB) {
    w.WriteHeader(http.StatusOK)
}
