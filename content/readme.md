# About

Welcome to the about section of the theme. 🎉

Here you can add valid markdown and have it rendered directly in the `/readme` page.

 - [Official Biography](https://nivenly.com/bio)
 - [Theme GitHub](https://github.com/kris-nova/prine)

### Code Samples

You can easily add code samples to the theme, and they will be rendered on the website.
This snippet is taken from my Kubernetes application management tool called [naml](https://github.com/kris-nova/naml)

```go
// Deployable is an interface that can be implemented
// for deployable applications.
type Deployable interface {

    // Install will attempt to install in Kubernetes
    Install(client kubernetes.Interface) error

    // Uninstall will attempt to uninstall in Kubernetes
    Uninstall(client kubernetes.Interface) error

    // Meta returns a NAML Meta structure which embed Kubernetes *metav1.ObjectMeta
    Meta() *AppMeta

    // Objects will return the runtime objects defined for each application
    Objects() []runtime.Object
}
```

Make sure you valid Syntax highlighting code blocks.

You can add lists, images, and shortcodes here directly in the markdown.
