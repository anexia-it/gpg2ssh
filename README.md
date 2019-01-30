# gpg2ssh

`gpg2ssh` is a tiny utility which allows conversion of an OpenPGP (public) key
to an SSH public key, usable in `authorized_keys`.

## install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/anexia-it/gpg2ssh
```

## usage

```sh
gpg2ssh armored_gpg_key.asc >/tmp/ssh_key.pub
```
