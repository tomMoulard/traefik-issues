# Reply

Hi @FlomoN,

Thanks for your interest in Traefik!

Sorry, I can't really understand your issue here.

To me, `X-Forwarded-Host` is used to give the backend service an information about the client query.

> The X-Forwarded-Host (XFH) header is a de-facto standard header for identifying the original host requested by the client in the Host HTTP request header.
> ~ [MDN Web Doc](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host)

Thus traefik will only give the host to the concerned service, not the auth service.

![](https://doc.traefik.io/traefik/v2.4/assets/img/middleware/authforward.png)

cf [traefik doc](https://doc.traefik.io/traefik/v2.4/middlewares/forwardauth/)

And traefik will create a new request to auth the user.

Can you please explain what you are trying to achieve ?

If you are trying to manage a Oauth proxy, take a look on how [authelia](https://www.authelia.com/) [did](https://www.authelia.com/docs/deployment/supported-proxies/traefik2.x.html#basic-authentication).

# Reply
I think you missed a point on authelia configuration, go take a look at the doc and the [community forum](https://community.traefik.io/search?q=authelia) ([hint](https://github.com/tomMoulard/make-my-server/tree/authelia/authelia#adding-authelia-to-another-service)).

As for the `trustForwardHeader` option, the [doc](https://doc.traefik.io/traefik/middlewares/forwardauth/#trustforwardheader) states:
> Set the trustForwardHeader option to true to trust all X-Forwarded-* headers.

Which says that traefik will forward all `X-Forwarded-*` headers if the option value is set to `true`.

To me you are stating a configuration error, and I think this discution should be in the [community traefik forum](https://community.traefik.io/).

