package dictionary

import (
	"testing"
)

//测试Search方法，参数为map和key
func TestSearchDictionary(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test"}
	got := Search(dictionary, "test")
	want := "this is a test"
	assertString(t, got, want)
}

//测试Search方法，Search添加Dictionary Type作为Reciever
func TestSearch(t *testing.T) {
	//Happy Path测试
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("finding key:test", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertString(t, got, want)
	})
	//测试出错的情况
	t.Run("finding an unknown key:wtf", func(t *testing.T) {
		_, err := dictionary.Search("wtf")
		want := errorNotFound
		assertError(t, err, want)
	})
}

//测试Add方法，Add添加Dictionary作为Reciever
func TestAdd(t *testing.T) {
	//happy path测试
	t.Run("add a word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		want := "this is just a test"
		dictionary.Add(key, want)
		assertDefinition(t, dictionary, key, want)
	})
	//测试出错的情况，add插入的key已存在
	t.Run("add an existing word", func(t *testing.T) {
		key := "test"
		definition := "this is a new test"
		dictionary := Dictionary{key: definition}
		err := dictionary.Add(key, "new test")
		assertError(t, err, errorKeyExisting)
		assertDefinition(t, dictionary, key, definition)
	})
}

//测试Update方法
func TestUpdate(t *testing.T) {
	//happy path路径测试
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is a new test"
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	//更新操作遇到新key
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "this is just a test"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, errKeyNotExisting)
	})
}

//测试Delete方法
func TestDelete(t *testing.T) {
	//happy path路径测试
	t.Run("delete a existed word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, errorNotFound)
	})
	//删除一个不存在的word
	t.Run("delete a not existed word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)
		assertError(t, err, errKeyNotExisting)
	})
}

//工具方法
func assertDefinition(t *testing.T, dictionary Dictionary, key, want string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find an error:", err)
	}
	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

//工具方法
func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("error should not be nil")
	}
	if got != want {
		t.Errorf("got %s, want %s", got.Error(), want.Error())
	}
}

//工具方法
func assertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
