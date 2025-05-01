/* Most Active Authors
 * In this challenge, the REST API contains information about a collection of users and articles they created.  Given the threshold value, the goal is to use the API to get the list of most active authors.  Specifically, the list of usernames of the users with submission count strictly greater than the given threshold.  The list of usernames must be returned in the order the users appear in the results.
 *
 * To access the collection of users perform HTTP GET request to:
 * https://jsonmock.hackerrank.com/api/article_users?page=<pageNumber>
 * where <pageNumber> is an integer denoting the page of the results to return.
 *
 * For example, GET request to:
 * https://jsonmock.hackerrank.com/api/article_users/search?page=2
 * will return the second page of the collection of users.  Pages are numbered from 1, so in order to access the first page, you need to ask for page numer 1.
 */
package main

import (
	"fmt"
	"net/http"
	"time"
)

// - Query REST API
//   - Iterate over all pages (1..n)
// - Process API response JSON
//   - filter data for users above threshold
//     - while we don't need data now keep all data for future
//   - submission count greater than threshold
//     ! time values in JSON are inconsistent format
// - Return the usernames
//   - in order in which the were received (append)

/* example content

{
  "page": 1,
  "per_page": 10,
  "total": 15,
  "total_pages": 2,
  "data": [
    {
      "id": 1,
      "username": "epaga",
      "about": "Java developer / team leader at inetsoftware.de by day\u003Cp\u003EiOS developer by night\u003Cp\u003Ehttp://www.mindscopeapp.com\u003Cp\u003Ehttp://inflightassistant.info\u003Cp\u003Ehttp://appstore.com/johngoering\u003Cp\u003E[ my public key: https://keybase.io/johngoering; my proof: https://keybase.io/johngoering/sigs/I1UIk7t3PjfB5v2GI-fhiOMvdkzn370_Z2iU5GitXa0 ]\u003Cp\u003Ehnchat:oYwa7PJ4Yaf1Vw9Om4ju",
      "submitted": 654,
      "updated_at": "2019-08-29T13:45:12.000Z",
      "submission_count": 197,
      "comment_count": 439,
      "created_at": 1301039509
    },
*/

type PageData struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       UserData `json:"data"`
}

type UserData struct {
	Id              int       `json:"id"`
	Username        string    `json:"username"`
	About           string    `json:"about"`
	Submitted       int       `json:"submitted"`
	UpdatedAt       time.Time `json:"updated_at"`
	SubmissionCount int       `json:"submission_count"`
	CommentCount    int       `json:"comment_count"`
	CreatedAt       time.Time `json:"created_at"`
}

type URL string

func getPageData(url *URL, page string, data *PageData) {
	r, err := http.Get(*url)
	if err != nil {
		fmt.Println("could not connect to", *url, err)
	}
}

func main() {

}
