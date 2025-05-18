package model

/**

  原型类的通用的接口，
  这样即使不知道对象具体类的情况下也能复制对象

*/

type ProtoInterface interface {
    Clone() ProtoInterface
}
