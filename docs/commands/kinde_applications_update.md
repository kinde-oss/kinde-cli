## kinde applications update

Update Application

### Synopsis

Updates a client's settings. For more information, read [Applications in
Kinde](https://docs.kinde.com/build/applications/about-applications)

Required scope: update:applications

API: PATCH /api/v1/applications/{application_id}

```
kinde applications update [flags]
```

### Options

```
      --application-id string               The identifier for the application (required)
      --body string                         Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --has-marketing-consent               Show a marketing consent checkbox on the sign-up page
      --has-register-link-on-sign-in-page   Allow users to switch to the register page from the sign-in page
      --has-sign-in-link-on-sign-up-page    Allow users to switch to the login page from the sign-up page
      --has-sign-in-with-sso-button         When home realm discovery is configured, users see a button to prompt them to use their work email
  -h, --help                                help for update
      --homepage-uri string                 The homepage link to your application
      --is-access-control-enabled           Whether role-based access control is enforced for the application
      --is-allow-faceless-auth              Bypass Kinde's sign up and sign in screens and use your own design
      --is-ask-for-name                     Show fields to collect name details from users signing up with email or phone
      --language-key string                 The application's language key
      --login-uri string                    The default login route for resolving session issues
      --logout-uris strings                 The application's logout uris (comma-separated list)
      --name string                         The application's name
      --redirect-uris strings               The application's redirect uris (comma-separated list)
      --use-gravatar-fallback               Use a backup image if a profile picture is not available
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

