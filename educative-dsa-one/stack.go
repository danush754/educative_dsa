package educativedsa

const LIMIT = 4

type Stack struct {
	ix   int
	data [LIMIT]int
}

func CreateStack() *Stack {
	var newStack Stack

	return &newStack
}

func (st *Stack) Push(n int) {

	if st.data[st.ix] == LIMIT {
		return
	}

	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	if st.ix > 0 {

		st.ix--
		element := st.data[st.ix]
		st.data[st.ix] = 0
		return element
	}

	return -1
}
