package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"torelog_bygolang/app/models"
	"torelog_bygolang/config"
)

// view周りの共通化ファイルをまとめる
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// クッキーを取得する
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("無効なセッション")
		}
	}

	return sess, err
}

// サーバーの設定
func StartMainServer() (err error) {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/trainingLog", index)
	http.HandleFunc("/logout", logout)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
