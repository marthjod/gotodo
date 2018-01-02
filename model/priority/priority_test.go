package priority

import "testing"

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
