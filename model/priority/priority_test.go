package priority

import (
	"reflect"
	"testing"
)

var prios = map[string]Priority{
	"(A)": A,
	"(B)": B,
	"(C)": C,
	"(D)": D,
	"(E)": E,
	"(F)": F,
	"(G)": G,
	"(H)": H,
	"(I)": I,
	"(J)": J,
	"(K)": K,
	"(L)": L,
	"(M)": M,
	"(N)": N,
	"(O)": O,
	"(P)": P,
	"(Q)": Q,
	"(R)": R,
	"(S)": S,
	"(T)": T,
	"(U)": U,
	"(V)": V,
	"(W)": W,
	"(X)": X,
	"(Y)": Y,
	"(Z)": Z,
	"(Ä)": None,
}

func TestGetPriority(t *testing.T) {
	for k, v := range prios {
		if GetPriority(k) != v {
			t.Errorf("Expected %q, but got %q", v, GetPriority(k))
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	var expected = []struct {
		in  Priority
		out []byte
	}{
		{
			in:  A,
			out: []byte(`"A"`),
		},
		{
			in:  None,
			out: []byte(`null`),
		},
	}

	for _, exp := range expected {
		actual, err := exp.in.MarshalJSON()
		if err != nil {
			t.Fatal(err.Error())
		}

		if !reflect.DeepEqual(actual, exp.out) {
			t.Errorf("Expected %s, got %s", exp.out, actual)
		}
	}
}
