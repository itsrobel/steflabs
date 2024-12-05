// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

// TODO: update colors to use theme
func NavLinks() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"/\"><h1>Home</h1></a> <a href=\"/about\"><h1>About</h1></a> <span data-toggle-theme=\"dim,wireframe\" data-act-class=\"pl-4\" class=\"border rounded-full border-green-700 flex items-center cursor-pointer w-10 transition-all duration-300 ease-in-out pl-0\"><span class=\"rounded-full w-3 h-3 m-1 bg-green-700\"></span></span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Nav() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><div class=\"hidden md:flex w-full md:justify-center\"><nav class=\" navbar bg-base-100 flex-none shadow-2xl font-bold font-excali flex justify-center space-x-8 mt-8 rounded-lg md:w-4/5 \">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavLinks().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</nav></div><!-- Full-Screen Mobile Menu --><div class=\"mobile-menu fixed inset-0 dark:bg-gray-800 bg-gray-200 bg-opacity-95 hidden flex flex-col items-center justify-center space-y-4 z-50 transform -translate-x-full opacity-0 transition-all duration-500 ease-in-out\"><button class=\"close-button absolute top-4 right-4 focus:outline-none\"><i class=\"fa-solid fa-arrow-right-from-bracket\"></i></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavLinks().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"md:hidden\"><div class=\"z-10 fixed bottom-4 right-4\"><button class=\"menu-button bg-blue-500 hover:bg-blue-700 text-white font-bold w-10 h-10 rounded-full shadow-lg focus:outline-none\"><i class=\"fa-solid fa-cube\"></i></button></div></div></div><script>\n  document.addEventListener('DOMContentLoaded', function () {\n    const btn = document.querySelector('button.menu-button');\n    const closeBtn = document.querySelector('button.close-button');\n    const menu = document.querySelector('.mobile-menu');\n\n    btn.addEventListener('click', () => {\n      menu.classList.remove('hidden');\n      setTimeout(() => {\n        menu.classList.add('translate-x-0', 'opacity-100');\n        menu.classList.remove('-translate-x-full', 'opacity-0');\n      }, 10); // Small delay to ensure transition works\n    });\n\n    closeBtn.addEventListener('click', () => {\n      menu.classList.remove('translate-x-0', 'opacity-100');\n      menu.classList.add('-translate-x-full', 'opacity-0');\n      setTimeout(() => {\n        menu.classList.add('hidden');\n      }, 500); // Match duration with CSS transition\n    });\n  });\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate