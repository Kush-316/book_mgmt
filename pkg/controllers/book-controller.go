package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/Kush-316/book_mgmt/pkg/utils"
	"github.com/Kush-316/book_mgmt/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res,_ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_ := models.GetBookById(ID)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_ := models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication !=""{
		bookDetails.Publication = updateBook.Publication
	}
	models.UpdateBookDetails(bookDetails)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}