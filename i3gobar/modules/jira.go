package modules

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/abibby/i3/i3gobar/icon"
	"github.com/davecgh/go-spew/spew"
)

type JiraClient struct {
	client  *http.Client
	baseUrl string
}

func (c *JiraClient) fetch(method, url string, modifyRequest func(r *http.Request)) error {
	body := bytes.NewBufferString(`{"client_id":"tDP5by46cc3gEck7d2vbHZsqsfrDK6t9","username":"abibby@mero.co","password":"x8h5ct6SKWvifa8nN85vgGCnP","realm":"eyJhcHBsaWNhdGlvbktleSI6ImppcmEifQ==","credential_type":"http://auth0.com/oauth/grant-type/password-realm"}`)
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	if modifyRequest != nil {
		modifyRequest(r)
	}

	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
	fmt.Println()

	return nil
}
func (c *JiraClient) Post(url string) error {
	return c.fetch("POST", url, nil)
}
func (c *JiraClient) Get(url string) error {
	return c.fetch("GET", url, nil)
}

func openJiraClient() (*JiraClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	// token, err := pass("atlassian.com")
	// if err != nil {
	// 	return nil, err
	// }

	client := &http.Client{
		Jar: jar,
	}

	c := &JiraClient{
		client: client,
	}

	// https://id.atlassian.com/login
	// html>body[data-app-state] has client id and stuff
	/*
		{
			"appConfig": {
				"contextPath": "",
				"auth0Config": {
					"clientId": "tDP5by46cc3gEck7d2vbHZsqsfrDK6t9",
					"tenant": "atlassian-account-prod",
					"domain": "auth.atlassian.com",
					"tokenIssuer": "https://atlassian-account-prod.pus2.auth0.com",
					"callbackUrl": "https://id.atlassian.com/login/callback"
				},
				"recaptchaKeySite": "6LewHQcTAAAAAJgaYVKlQOahz4gnQME8wqUA0z0J",
				"segmentIoKey": "cb2egpwag7",
				"recaptchaInvisibleKeySite": "6LcqAHAUAAAAAKcO583Ymvnq-uRBDPq4njcoW-jK",
				"castleAppId": "337683121243755",
				"bitbucketSignupUrlOverrideEnabled": false,
				"sentryUrl": "https://71e54c28be0d49f0bcd732ab30f35faa@sentry.io/275199",
				"recaptchaEnable": true,
				"bitbucketSignupUrl": "https://bitbucket.org/account/signup",
				"marketingConsentApiUrl": "https://preferences.atlassian.com/rest",
				"googleAuthClientId": "596149463257-9oquqfivs9on8t8erq23c8qso6vk3cp1.apps.googleusercontent.com"
			},
			"featureFlags": {
				"aid_signup.microsoft.auth.enabled": true,
				"aid_signup.authenticate.via.id.authentication": false,
				"aid.sign_with_asap.verify_email_link.enabled": false,
				"aid.sign_with_asap.welcome_link.enabled": false,
				"aid_signup.apple.auth.enabled": true,
				"id-authentication.IDZERO-382-login-attempts-limit.enabled": true,
				"aid_signup.email_reverification.social_login.enabled": false,
				"aid_signup.domain.claim.data.signup.analytics": true,
				"aid_signup.migrating.users.redirect.to.welcome": false,
				"aid_signup.remove.continue.mobile.verify.email.action": false,
				"aid_signup.google.callback.adg3": true,
				"sign-in-with-slack.enabled": false,
				"aid_signup.experiment.user_segmentation": "variation",
				"aid_signup.email_reverification.existing_flows.enabled": false,
				"aid_signup.passwordless.signup": false,
				"aid_signup.disallow.passwordless.login.for.google.users": true,
				"aid_signup.apple.show.hidden.email.warning.enabled": true,
				"aid.sign_with_asap.verification_email_sent_page.enabled": false,
				"aid_signup.sev.enabled": true,
				"aid_signup.default.continue.skip.profile.action": true,
				"aid_signup.domain.claim.data.login.analytics": true,
				"aid.sign_with_asap.welcome_email_sent_page.enabled": false,
				"aid_signup.forced.redirect.session.distribution.enabled": true,
				"aid_signup.csrf.refresh.align.with.sis": true,
				"aid_signup.bans.validate.authentication": true,
				"aid_signup.shadow.call.to.id.authentication": true
			},
			"csrfToken": "b60d9e3ae1daa769039b995d8df7118ca3969bcd",
			"microbranding": {
				"application": "jira",
				"applicationNameShort": "jira.atlassian.com",
				"applicationLogoClass": "jira",
				"isEmbedded": "false",
				"applicationName": "Atlassian Bug Reporting & Feature Requests",
				"applicationBaseURI": "http://idm-web-staging.private.atlassian.com:8080/jira/"
			},
			"tenantCloudId": "ee838988-cd60-4563-a11d-5a96a378969a"
		}
	*/
	err = c.fetch("POST", "https://auth.atlassian.com/co/authenticate", func(r *http.Request) {
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Origin", "https://id.atlassian.com")
		r.Header.Add("Cache-Control", "no-cache")
	})
	if err != nil {
		return nil, err
	}

	c.Get("https://auth.atlassian.com/authorize?client_id=tDP5by46cc3gEck7d2vbHZsqsfrDK6t9&redirect_uri=https%3A%2F%2Fid.atlassian.com%2Flogin%2Fcallback%3Fapplication%3Djira%26continue%3Dhttps%253A%252F%252Fmerotechnologies.atlassian.net%252Flogin%253FredirectCount%253D1%2526application%253Djira%26email%3Dabibby%2540mero.co%26errorCode&response_type=code&state=eyJjc3JmVG9rZW4iOiJkNzMwZjk0MjI5MmUzODM4YzRkNGY1Y2IxN2E2NjY3MDBjNTY1NTkxIiwiYW5vbnltb3VzSWQiOiJjZWE5NWI3NS05YzBlLTQyZjQtYmM4Mi0wZmUxYmU3NmEwNTMifQ%3D%3D&realm=eyJhcHBsaWNhdGlvbktleSI6ImppcmEifQ%3D%3D&scope=openid%20profile&login_ticket=goyQFY8sYUfP58GoeXa80AZgcdmQScna&auth0Client=eyJuYW1lIjoiYXV0aDAuanMiLCJ2ZXJzaW9uIjoiOS4xMy4yIn0%3D")

	// err = c.Get("https://id.atlassian.com/login?application=jira&continue=https%3A//merotechnologies.atlassian.net/login?redirectCount%3D1%26application%3Djira&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJzaWdudXAiLCJpYXQiOjE2MDUxODMxNjYsImV4cCI6MTYwNTE4MzI4NiwidXNlcklkIjoiNWVjNmM1NzhiMTk5YWEwYzEzZjk5ODRhIiwibG9naW5UeXBlIjoicGFzc3dvcmRMb2dpbiIsIm1hcmtlZFZlcmlmaWVkIjoiZmFsc2UiLCJzY29wZSI6IkxvZ2luIn0.zVHHqc08K4BQclCILxd0vPj1PmtiUzKXwjKfElVJ7aI")
	// if err != nil {
	// 	return nil, err
	// }

	// spew.Dump(jar)
	// os.Exit(1)

	return c, nil
}

func JiraNotifications() string {
	client, err := openJiraClient()
	if err != nil {
		return ""
	}

	err = client.Get("https://merotechnologies.atlassian.net/gateway/api/notification-log/api/2/notifications?cloudId=ee838988-cd60-4563-a11d-5a96a378969a&direct=true&includeContent=true&limit=8")
	if err != nil {
		spew.Dump(err)
		return ""
	}
	return fmt.Sprintf("%s %d", icon.Jira, 0)
}
