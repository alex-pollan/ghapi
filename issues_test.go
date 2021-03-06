package ghapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const getIssue1347Response string = `{
  "id": 1,
  "url": "https://api.github.com/repos/octocat/Hello-World/issues/1347",
  "labels_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/labels{/name}",
  "comments_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/comments",
  "events_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/events",
  "html_url": "https://github.com/octocat/Hello-World/issues/1347",
  "number": 1347,
  "state": "open",
  "title": "Found a bug",
  "body": "I'm having a problem with this.",
  "user": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "labels": [
    {
      "url": "https://api.github.com/repos/octocat/Hello-World/labels/bug",
      "name": "bug",
      "color": "f29513"
    }
  ],
  "assignee": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "milestone": {
    "url": "https://api.github.com/repos/octocat/Hello-World/milestones/1",
    "html_url": "https://github.com/octocat/Hello-World/milestones/v1.0",
    "labels_url": "https://api.github.com/repos/octocat/Hello-World/milestones/1/labels",
    "id": 1002604,
    "number": 1,
    "state": "open",
    "title": "v1.0",
    "description": "Tracking milestone for version 1.0",
    "creator": {
      "login": "octocat",
      "id": 1,
      "avatar_url": "https://github.com/images/error/octocat_happy.gif",
      "gravatar_id": "",
      "url": "https://api.github.com/users/octocat",
      "html_url": "https://github.com/octocat",
      "followers_url": "https://api.github.com/users/octocat/followers",
      "following_url": "https://api.github.com/users/octocat/following{/other_user}",
      "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
      "organizations_url": "https://api.github.com/users/octocat/orgs",
      "repos_url": "https://api.github.com/users/octocat/repos",
      "events_url": "https://api.github.com/users/octocat/events{/privacy}",
      "received_events_url": "https://api.github.com/users/octocat/received_events",
      "type": "User",
      "site_admin": false
    },
    "open_issues": 4,
    "closed_issues": 8,
    "created_at": "2011-04-10T20:09:31Z",
    "updated_at": "2014-03-03T18:58:10Z",
    "closed_at": "2013-02-12T13:22:01Z",
    "due_on": "2012-10-09T23:39:01Z"
  },
  "locked": false,
  "comments": 0,
  "pull_request": {
    "url": "https://api.github.com/repos/octocat/Hello-World/pulls/1347",
    "html_url": "https://github.com/octocat/Hello-World/pull/1347",
    "diff_url": "https://github.com/octocat/Hello-World/pull/1347.diff",
    "patch_url": "https://github.com/octocat/Hello-World/pull/1347.patch"
  },
  "closed_at": null,
  "created_at": "2011-04-22T13:33:48Z",
  "updated_at": "2011-04-22T13:33:48Z",
  "closed_by": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  }
}`

const getIssueComment1Response string = `{
  "id": 1,
  "url": "https://api.github.com/repos/octocat/Hello-World/issues/comments/1",
  "html_url": "https://github.com/octocat/Hello-World/issues/1347#issuecomment-1",
  "body": "Me too",
  "user": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "created_at": "2011-04-14T16:00:49Z",
  "updated_at": "2011-04-14T16:00:49Z"
}`

func expectIssue1347(t *testing.T, issue *IssueResponse) {
	expectNotNil(t, issue, "issue")
	expect(t, 1, issue.ID, "issue.Id")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347", issue.URL, "issue.Url")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347/labels{/name}", issue.LabelsURL, "issue.LabelsUrl")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347/comments", issue.CommentsURL, "issue.CommentsUrl")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347/events", issue.EventsURL, "issue.EventsUrl")
	expect(t, "https://github.com/octocat/Hello-World/issues/1347", issue.HTMLURL, "issue.HtmlUrl")
	expect(t, 1347, issue.Number, "issue.Number")
	expect(t, "open", issue.State, "issue.State")
	expect(t, "Found a bug", issue.Title, "issue.Title")
	expect(t, "I'm having a problem with this.", issue.Body, "issue.Body")

	expectOctocatUser(t, &issue.User, "issue.User")

	expect(t, 1, len(issue.Labels), "len(issue.Labels)")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/labels/bug", issue.Labels[0].URL, "issue.Labels[0].Url")
	expect(t, "bug", issue.Labels[0].Name, "issue.Labels[0].Name")
	expect(t, "f29513", issue.Labels[0].Color, "issue.Labels[0].Color")

	expectOctocatUser(t, issue.Assignee, "issue.Assignee")

	expectNotNil(t, issue.Milestone, "issue.Milestone")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/milestones/1", issue.Milestone.URL, "issue.Milestone.Url")
	expect(t, "https://github.com/octocat/Hello-World/milestones/v1.0", issue.Milestone.HTMLURL, "issue.Milestone.HtmlUrl")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/milestones/1/labels", issue.Milestone.LabelsURL, "issue.Milestone.LabelsUrl")
	expect(t, 1002604, issue.Milestone.ID, "issue.Milestone.Id")
	expect(t, 1, issue.Milestone.Number, "issue.Milestone.Number")
	expect(t, "open", issue.Milestone.State, "issue.Milestone.State")
	expect(t, "v1.0", issue.Milestone.Title, "issue.Milestone.Title")
	expect(t, "Tracking milestone for version 1.0", issue.Milestone.Description, "issue.Milestone.Description")
	expectOctocatUser(t, &issue.Milestone.Creator, "issue.Milestone.Creator")
	expect(t, 4, issue.Milestone.OpenIssues, "issue.Milestone.OpenIssues")
	expect(t, 8, issue.Milestone.ClosedIssues, "issue.Milestone.ClosedIssues")
	expect(t, date("2011-04-10T20:09:31Z"), issue.Milestone.CreatedAt, "issue.Milestone.CreatedAt")
	expect(t, date("2014-03-03T18:58:10Z"), issue.Milestone.UpdatedAt, "issue.Milestone.UpdatedAt")
	expect(t, date("2013-02-12T13:22:01Z"), issue.Milestone.ClosedAt, "issue.Milestone.ClosedAt")
	expect(t, date("2012-10-09T23:39:01Z"), issue.Milestone.DueOn, "issue.Milestone.DueOn")

	expect(t, false, issue.Locked, "issue.Locked")
	expect(t, 0, issue.Comments, "issue.Comments")

	expectNotNil(t, issue.PullRequest, "issue.PullRequest")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/pulls/1347", issue.PullRequest.URL, "issue.PullRequest.Url")
	expect(t, "https://github.com/octocat/Hello-World/pull/1347", issue.PullRequest.HTMLURL, "issue.PullRequest.HtmlUrl")
	expect(t, "https://github.com/octocat/Hello-World/pull/1347.diff", issue.PullRequest.DiffURL, "issue.PullRequest.DiffUrl")
	expect(t, "https://github.com/octocat/Hello-World/pull/1347.patch", issue.PullRequest.PatchURL, "issue.PullRequest.PatchUrl")

	expectNil(t, issue.ClosedAt, "issue.ClosedAt")
	expect(t, date("2011-04-22T13:33:48Z"), issue.CreatedAt, "issue.CreatedAt")
	expect(t, date("2011-04-22T13:33:48Z"), issue.UpdatedAt, "issue.UpdatedAt")

	expectOctocatUser(t, issue.ClosedBy, "issue.ClosedBy")
}

func expectOctocatUser(t *testing.T, u *User, prefix string) {
	expectNotNil(t, u, prefix)

	expect(t, "octocat", u.Login, prefix+".Login")
	expect(t, 1, u.ID, prefix+".Id")
	expect(t, "https://github.com/images/error/octocat_happy.gif", u.AvatarURL, prefix+".AvatarUrl")
	expect(t, "", u.GravatarID, prefix+".GravatarId")
	expect(t, "https://api.github.com/users/octocat", u.URL, prefix+".Url")
	expect(t, "https://github.com/octocat", u.HTMLURL, prefix+".HtmlUrl")
	expect(t, "https://api.github.com/users/octocat/followers", u.FollowersURL, prefix+".FollowersUrl")
	expect(t, "https://api.github.com/users/octocat/following{/other_user}", u.FollowingURL, prefix+".FollowingUrl")
	expect(t, "https://api.github.com/users/octocat/gists{/gist_id}", u.GistsURL, prefix+".GistsUrl")
	expect(t, "https://api.github.com/users/octocat/starred{/owner}{/repo}", u.StarredURL, prefix+".StarredUrl")
	expect(t, "https://api.github.com/users/octocat/subscriptions", u.SubscriptionsURL, prefix+".SubscriptionsUrl")
	expect(t, "https://api.github.com/users/octocat/orgs", u.OrganizationsURL, prefix+".OrganizationsUrl")
	expect(t, "https://api.github.com/users/octocat/repos", u.ReposURL, prefix+".ReposUrl")
	expect(t, "https://api.github.com/users/octocat/events{/privacy}", u.EventsURL, prefix+".EventsUrl")
	expect(t, "https://api.github.com/users/octocat/received_events", u.ReceivedEventsURL, prefix+".ReceivedEventsUrl")
	expect(t, "User", u.Type, prefix+".Type")
	expect(t, false, u.SiteAdmin, prefix+".SiteAdmin")
}

func TestIssueApi_DeleteIssueComment(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/comments/1" {

			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")
			expect(t, "DELETE", r.Method, "r.Method")
			expect(t, "", string(b), "r.Body")

			w.WriteHeader(200)
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	err := api.Issue.DeleteIssueComment(1)
	waitSignal(t, signal)

	expectNil(t, err, "err")
}

func TestIssueApi_DeleteIssueComment_404(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/comments/1" {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	err := api.Issue.DeleteIssueComment(2)
	waitSignal(t, signal)

	if err != nil {
		if e, ok := err.(*ErrHTTPError); !ok {
			t.Fatalf("err is not of type *ErrHttpError, is %T", err)
		} else {
			expect(t, 404, e.StatusCode, "e.StatusCode")
		}
	} else {
		t.Fatal("expected error")
	}
}

func TestIssueApi_GetIssue(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/1347" {
			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			_, err = w.Write([]byte(getIssue1347Response))

			expectNil(t, err, "err")
			expect(t, "GET", r.Method, "r.Method")
			expect(t, "", string(b), "r.Body")
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	issue, err := api.Issue.GetIssue(1347)
	waitSignal(t, signal)

	if err != nil {
		t.Fatal(err)
	}

	expectIssue1347(t, issue)
}

func TestIssueApi_GetIssue_ReturnsErrOnHttpErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/1347" {
			w.WriteHeader(500)
		} else {
			t.Fatalf("unexpected url %s", r.URL)
		}
	})
	defer ts.Close()

	issue, err := api.Issue.GetIssue(1347)
	waitSignal(t, signal)

	expectNil(t, issue, "issue")
	expectErrHTTPError500(t, err)
}

func TestIssueApi_GetIssue_ReturnsErrOnJsonDecodeErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/1347" {
			_, err := w.Write([]byte("junk" + getIssue1347Response))
			if err != nil {
				t.Fatal(err)
			}
		} else {
			t.Fatalf("unexpected url %s", r.URL)
		}
	})
	defer ts.Close()

	issue, err := api.Issue.GetIssue(1347)
	waitSignal(t, signal)

	expectNil(t, issue, "issue")
	expectJSONSyntaxError(t, err, "invalid character 'j' looking for beginning of value")
}

func TestIssueApi_UpdateIssueAssignee(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/1347" {

			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			_, err = w.Write([]byte(getIssue1347Response))

			expectNil(t, err, "err")
			expect(t, "PATCH", r.Method, "r.Method")
			expect(t, `{"assignee":"new-assignee"}`, string(b), "r.Body")
		} else {
			t.Fatalf("unexpected url %s", r.URL)
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	issue, err := api.Issue.UpdateIssueAssignee(1347, "new-assignee")
	waitSignal(t, signal)

	expectNil(t, err, "err")
	expectIssue1347(t, issue)
}

func TestIssueApi_UpdateIssueAssignee_ReturnsErrOnHttpErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/1347" {

			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			expect(t, "PATCH", r.Method, "r.Method")
			expect(t, `{"assignee":"new-assignee"}`, string(b), "r.Body")

			w.WriteHeader(500)
		} else {
			t.Fatalf("unexpected url %s", r.URL)
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	issue, err := api.Issue.UpdateIssueAssignee(1347, "new-assignee")
	waitSignal(t, signal)

	expectNotNil(t, err, "err")
	expectNil(t, issue, "issue")
	if err != nil {
		if e, ok := err.(*ErrHTTPError); !ok {
			t.Fatalf("err is not of type *ErrHttpError, is %T", err)
		} else {
			expect(t, 500, e.StatusCode, "e.StatusCode")
		}
	} else {
		t.Fatal("expected error")
	}
}

func TestIssueApi_UpdateIssueAssignee_ReturnsErrOnJsonDecodeErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/1347" {

			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			_, err = w.Write([]byte("junk" + getIssue1347Response))

			expectNil(t, err, "err")
			expect(t, "PATCH", r.Method, "r.Method")
			expect(t, `{"assignee":"new-assignee"}`, string(b), "r.Body")
		} else {
			t.Fatalf("unexpected url %s", r.URL)
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	issue, err := api.Issue.UpdateIssueAssignee(1347, "new-assignee")
	waitSignal(t, signal)

	expectNil(t, issue, "issue")
	expectJSONSyntaxError(t, err, "invalid character 'j' looking for beginning of value")
}

func TestIssueApi_updateIssueByUrl_ReturnsErrOnJsonMarshalErr(t *testing.T) {
	api := makeGitHubAPI()

	data := struct {
		Assignee func(string) `json:"assignee"`
	}{nil}

	issue, err := api.Issue.updateIssueByURL("http://error.org", data)

	expectNil(t, issue, "issue")
	if err != nil {
		if e, ok := err.(*json.UnsupportedTypeError); !ok {
			t.Fatalf("err is not of type *json.UnsupportedTypeError, is %T", err)
		} else {
			expect(t, "json: unsupported type: func(string)", e.Error(), "e.Error()")
		}
	} else {
		t.Fatal("expected error")
	}
}

func TestIssueApi_GetIssueComment(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/comments/1" {

			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			expect(t, "GET", r.Method, "r.Method")
			expectNotNil(t, b, "r.Body")
			expect(t, 0, len(b), "len(r.Body)")

			_, err = w.Write([]byte(getIssueComment1Response))
			expectNil(t, err, "err")
		} else {
			t.Fatalf("unexpected url %s", r.URL)
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	issueComment, err := api.Issue.GetIssueComment(1)
	waitSignal(t, signal)

	expectNil(t, err, "err")
	expectNotNil(t, issueComment, "issueComment")

	expect(t, 1, issueComment.ID, "issueComment.ID")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/comments/1", issueComment.URL, "issueComment.URL")
	expect(t, "https://github.com/octocat/Hello-World/issues/1347#issuecomment-1", issueComment.HTMLURL, "issueComment.HTMLURL")
	expect(t, "Me too", issueComment.Body, "issueComment.Body")
	expect(t, date("2011-04-14T16:00:49Z"), issueComment.CreatedAt, "issueComment.CreatedAt")
	expect(t, date("2011-04-14T16:00:49Z"), issueComment.UpdatedAt, "issueComment.UpdatedAt")
	expectOctocatUser(t, &issueComment.User, "issueComment.User")
}

func TestIssueApi_GetIssueComment_ReturnsErrOnHttpErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/comments/1" {

			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			expect(t, "GET", r.Method, "r.Method")
			expectNotNil(t, b, "r.Body")
			expect(t, 0, len(b), "len(r.Body)")

			w.WriteHeader(500)
		} else {
			t.Fatalf("unexpected url %s", r.URL)
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	issueComment, err := api.Issue.GetIssueComment(1)
	waitSignal(t, signal)

	expectNotNil(t, err, "err")
	expectNil(t, issueComment, "issueComment")
	if err != nil {
		if e, ok := err.(*ErrHTTPError); !ok {
			t.Fatalf("err is not of type *ErrHttpError, is %T", err)
		} else {
			expect(t, 500, e.StatusCode, "e.StatusCode")
		}
	} else {
		t.Fatal("expected error")
	}
}

func TestIssueApi_GetIssueComment_ReturnsErrOnJsonDecodeErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/comments/1" {

			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			expect(t, "GET", r.Method, "r.Method")
			expectNotNil(t, b, "r.Body")
			expect(t, 0, len(b), "len(r.Body)")

			_, err = w.Write([]byte("junk" + getIssueComment1Response))
			expectNil(t, err, "err")
		} else {
			t.Fatalf("unexpected url %s", r.URL)
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	issueComment, err := api.Issue.GetIssueComment(1)
	waitSignal(t, signal)

	expectNil(t, issueComment, "issueComment")
	expectJSONSyntaxError(t, err, "invalid character 'j' looking for beginning of value")
}
