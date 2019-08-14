// The terror package implements typed, user friendly errors in go. The typed errors parallel and are compatible with HTTP errors. A typed error has two basic components:
// * A user friendly error message meant for viewing by an end-user.
// * The original underlying error, if any.
//
// Implementations can thus easily support both communicating directly to the user while maintaining all the raw information one would want to capture in a log.
//
// terror also supports easy debug logging by either calling 'EchoErrorLog' from your test file or setting the `DEBUG_TERROR` environment variable to any non-blank value.
package terror
