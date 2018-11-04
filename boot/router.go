package boot

import (
    "gitee.com/johng/gf-home/app/ctl/document"
    "gitee.com/johng/gf/g/net/ghttp"
    "gitee.com/johng/gf/g"
)

// 统一路由注册.
func initRouter() {
    // 开发文档
    g.Server().BindHandler("/*path",    ctl_document.Index)
    g.Server().BindHandler("/hook",     ctl_document.UpdateHook)
    g.Server().BindHandler("/search",   ctl_document.Search)

    // 管理接口
    g.Server().EnableAdmin("/admin")

    // 某些浏览器会直接请求/favicon.ico文件，会产生404
    g.Server().BindHandler("/favicon.ico", func(r *ghttp.Request) {
        r.Response.ServeFile("/static/resource/image/favicon.ico")
    })

    // 为平滑重启管理页面设置HTTP Basic账号密码
    g.Server().BindHookHandler("/admin/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
        user := g.Config().GetString("doc.admin.user")
        pass := g.Config().GetString("doc.admin.pass")
        if !r.BasicAuth(user, pass) {
            r.Exit()
        }
    })

    // 强制跳转到HTTPS访问
    //g.Server().BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
    //    if !r.IsFileServe() && r.TLS == nil {
    //        r.Response.RedirectTo(fmt.Sprintf("https://%s%s", r.Host, r.URL.String()))
    //        r.Exit()
    //    }
    //})
}