## NextAuth + Custom Backend Authentication

This is an experimental project to implement custom backend authentication adapter for NextAuth (NextJS Authentication). The backend is developed by Golang + Go-Chi



> :exclamation:Warning: This is an experimental project, codebase is not structured for production grade system



**Experiment Target**

Experiment target is using a custom backend service to manage signup and login for social account as well as credential account.

* Before login a user must signup into that system using social account (google) or credential
* User must have to login with same provider after signup
  * If someone create account by using google account, then the user must have to use google account while logging into that system
  * If someone create account by using credential, then the user must have to use credential to logging into that account



**Setup**

* Google login setup
  * To enable google login use your own google OAuth credential in `next-web/.env` file
  * Set google `redirect_url` for both login and signup callback. This `google-signin` and `google-signup` values should be same as our `GoogleProvider` ids. Follow `next-web/pages/api/auth/[...nextauth].tsx` 
    *   `http://localhost:3000/api/auth/callback/google-signin`
    * `http://localhost:3000/api/auth/callback/google-signup`

* Email verification setup
  * To send email verification url to credential based registered user use your own email host configuration in `go-api/.env` file
