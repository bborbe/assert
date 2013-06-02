assert
======

assert for Go

Usage:
-----

    func TestSomething(t *testing.T) {
      result := 20 + 21
      err := assert.EqualsInt("number didn't match", 42, result)
      if err != nil {
        t.Error(err)
      }
    }

Available methods:
---------------

    func NotNullError(message string, value interface{}) (error) 
    
    func NullError(message string, value interface{}) (error)
    
    func EqualsString(message string, expected string, value string) (error)
    
    func EqualsInt(message string, expected int, value int) (error)
    
    func True(message string, value bool) (error)
    
    func False(message string, value bool) (error)