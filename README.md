<p align="center">
  <a href="https://github.com/wundergraph/custom-router-builds">
    <img src="img.png" width="500px" alt="gRPC Plugin Demo" />
  </a>
</p>

<p align="center">A template project for building your own custom Cosmo router with custom modules.</p>

<p align="center">
  <a href="https://cosmo-docs.wundergraph.com/router">Router Documentation</a> •
  <a href="https://cosmo-docs.wundergraph.com/router/custom-modules">Module Configuration</a>
</p>

## 📁 Examples

For an example of how to write custom modules, check out the [Router Examples project](https://github.com/wundergraph/router-examples), which contains an example, "myModule", of a custom router module.
This repository is designed to allow rapid, clean construction of custom routers with modules such that you can upgrade router versions simply by tracking our upstream and running `make`.

## 🚀 How to Use this Project
This project gives you a template for building a custom router with custom router modules. The "moduletemplate" package provides a basic implementation that can be loaded, documents how all of the hooks work, and does
nothing. It serves as a minimal framework for custom router modules where you just need to modify a few variables and add function code for your module's logic.

1. **Fork this repository** into your own. You can then use that forked repository as the basis for your router implementation. It's possible to have multiple forks that derive from a single "core" fork if, for instance, you
want to have different routers running different module configurations without adding any form of configuration management to `main.go`.
2. **Copy the moduletemplate** module for each new module you want to create. You can safely have multiple models, and the moduletemplate itself, so long as they each have different package name (per normal go standards).
3. **Modify the new modules** to have correct package names and to implement the hooks you need. For any hook you don't need, we recommend removing the function and the interface guard to improve performance and simplify
your code.
4. **Create appropriate configuration entries** in config.yaml for any configuration settings you added, along with documentation on how they work. This isn't strictly necessary, as the config.yaml used in the container will be
overridden by anything you provide on launch, but it's best practice to help people understand how your module is configured.
5. **Load your modules** by adding them to `main.go`. Each module is loaded using the `_ <fqdn-packagename>` syntax. The order you put them here doesn't matter; the `priority` variable you set in the module.go file determines load order.
6. **Run go mod tidy** to ensure all dependencies are up to date.
7. **Build the new router image** by running `make docker-build` or just `make`. By default, this will build a "cosmo-custom-router:latest" container image for linux/amd64. You can customize this by setting the `IMAGE_NAME`, `IMAGE_TAG`, `TARGETOS`, and `TARGETARCH` variables to build for a different platform, apply a different tag, or name the router image something different entirely.

### How the build works
The Makefile runs `go mod tidy` then triggers a docker build process, which is configured in the Dockerfile. The Dockerfile loads an appropriate golang layer for build, builds the router for the target architecture, then
creates a second stage that's distroless into which the statically linked router is placed.

### Updating the router
Whenever a new version of the router is released, this repository is updated to use that new version. Thus, to update the router, you need only track this repository as your upstream:

` git remote add upstream https://github.com/wundergraph/custom-router-builds.git `

Then when there's a new router release and you want to upgrade to it, presuming the branch you want to update to is `main`:

```
git fetch upstream
git checkout main 
git merge upstream/main
go mod tidy
make
```

You shouldn't run into merge conflicts if you've followed the instructions above and do not change this README.md file.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit examples, improvements, and bug fixes.

## License

Cosmo is licensed under the [Apache License, Version 2.0](LICENSE).
