/*
 * @Author: Alexleslie
 * @Date: 2022-04-01 01:11:25
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-04-01 01:14:34
 * @FilePath: \src\下载器\panel.go
 * @Description:
 */
// /*
//  * @Author: Alexleslie
//  * @Date: 2022-04-01 01:11:25
//  * @LastEditors: Alexleslie
//  * @LastEditTime: 2022-04-01 01:11:27
//  * @FilePath: \src\下载器\panel.go
//  * @Description:
//  */

package main

// import (
// 	"log"
// 	"os"

// 	"github.com/gotk3/gotk3/glib"
// 	"github.com/gotk3/gotk3/gtk"
// )

// func panel() {
// 	const appId = "com.nayoso.example"
// 	//每个gtk3程序都需要一步
// 	app, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)

// 	if err != nil {
// 		log.Fatal("Could not create application.", err)
// 	}

// 	//为activate事件绑定函数, activate会在程序启动时触发，也就是app.Run()时
// 	app.Connect("activate", func() {
// 		onActivate(app)
// 	})

// 	app.Run(os.Args) //运行gtkApplication
// }

// func onActivate(application *gtk.Application) {
// 	appWindow, err := gtk.ApplicationWindowNew(application) //创建window控件

// 	if err != nil {
// 		log.Fatal("Could not create application window.", err)
// 	}
// 	//设置窗口属性
// 	appWindow.SetTitle("Basic Application.")
// 	appWindow.SetDefaultSize(400, 400)
// 	//显示窗口
// 	appWindow.Show()
// }
