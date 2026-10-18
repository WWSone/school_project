// Package router — регистрация маршрутов HTTP.
//
// Здесь собирается http.ServeMux (или свой роутер) и связываются
// URL-пути REST API с хендлерами из пакета handlers, а также навешивается
// middleware. Реализацию пиши сам.
package router
