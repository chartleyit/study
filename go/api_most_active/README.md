# Most Active Authors

---

In this challenge, the REST API contains information about a collection of users and articles they created.  Given the threshold value, the goal is to use the API to get the list of most active authors.  Specifically, the list of usernames of the users with submission count strictly greater than the given threshold.  The list of usernames must be returned in the order the users appear in the results.

To access the collection of users perform HTTP GET request to:
https://jsonmock.hackerrank.com/api/article_users?page=<pageNumber>
where <pageNumber> is an integer denoting the page of the results to return.

For example, GET request to:
https://jsonmock.hackerrank.com/api/article_users/search?page=2
will return the second page of the collection of users.  Pages are numbered from 1, so in order to access the first page, you need to ask for page numer 1.