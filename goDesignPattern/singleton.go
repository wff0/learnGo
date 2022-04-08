package goDesignPattern

type Singleton struct {
}

//
//var singleton *Singleton
//
//func init() {
//	singleton := &Singleton{}
//}
//
//func getInstance() *Singleton {
//	return singleton
//}

//var (
//	lazyInstance *Singleton
//	once         = &sync.Once{}
//)
//
//func getLazyInstance() *Singleton {
//	if lazyInstance == nil {
//		once.Do(func() {
//			lazyInstance = &Singleton{}
//		})
//	}
//	return lazyInstance
//}
