# deploy-azure — docs

**Azure deploy.** Deploy to Azure Container Apps via the authenticated `az` CLI.

## Install

```bash
togo install togo-framework/deploy-azure
```

Registers on the [`deploy`](https://github.com/togo-framework/deploy) base; select it with **deploy.provider in togo.yaml (or DEPLOY_PROVIDER)**, then use **`togo deploy`**.

## Interface

`Deployer` — `Provision`/`Deploy`/`Destroy`/`Status` over a `Spec{App,Dir,BuildCmd,Host,User,Image,Region,Domain}` built from your `togo.yaml`.

## Usage & notes

Requires the `az` CLI installed + logged in. Deploys `spec.Image` to a Container App.

## Example

```bash
togo deploy --provider azure --dry-run   # preview the plan
togo deploy --provider azure
```

## Links

- [Azure Container Apps](https://learn.microsoft.com/azure/container-apps/)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/deploy-azure)
