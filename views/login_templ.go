// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/MSkrzypietz/rss/views/components"

func Login() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body><main class=\"min-h-screen w-full\"><form hx-post=\"/login\"><div class=\"flex items-center justify-center h-screen bg-background\"><div class=\"w-full max-w-md rounded-lg border bg-card text-card-foreground shadow-sm\"><div class=\"space-y-1 text-center flex flex-col space-y-1.5 p-6\"><div class=\"text-3xl font-bold leading-none tracking-tight\">Welcome Back</div><div class=\"text-sm text-muted-foreground\">Enter your API key to access your RSS feed.</div></div><div class=\"space-y-4 p-6 pt-0\"><div class=\"space-y-2\"><label for=\"apiKey\">API Key</label> <input id=\"apiKey\" name=\"apiKey\" type=\"text\" placeholder=\"Enter your API key\" class=\"w-full focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-stone-950 rounded border h-10 px-4 py-2\"></div></div><div class=\"flex items-center p-6 pt-0\"><button type=\"submit\" class=\"w-full bg-stone-950 text-white rounded h-10 px-4 py-2\">Log In</button></div></div></div></form></main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
