package discordbotsgg

import "testing"

func TestQueryParameters_String(t *testing.T) {
	const expected = "authorId=1&authorName=test&lib=test&limit=1&order=DESC&page=1&q=test&sort=username&unverified=true"

	queryParameters := &QueryParameters{
		Q:          "test",
		Page:       1,
		Limit:      1,
		AuthorID:   1,
		AuthorName: "test",
		Unverified: true,
		Lib:        "test",
		Sort:       "username",
		Order:      "DESC",
	}

	actual := queryParameters.String()

	if actual != expected {
		t.Errorf(
			"Unexpected QueryParameters string.\n\tGot: %s\n\tExpected: %s",
			actual,
			expected,
		)
	}
}