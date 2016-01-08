package main

import (
)

func permutations(iterable []*Person, r int) chan []*Person {
  ch := make(chan []*Person)

  go func() {
    pool := iterable
  	n := len(pool)

  	if r > n {
  		return
  	}

  	indices := make([]int, n)
  	for i := range indices {
  		indices[i] = i
  	}

  	cycles := make([]int, r)
  	for i := range cycles {
  		cycles[i] = n - i
  	}

  	result := make([]*Person, r)
  	for i, el := range indices[:r] {
  		result[i] = pool[el]
  	}

  	ch <- result

  	for n > 0 {
  		i := r - 1
  		for ; i >= 0; i -= 1 {
  			cycles[i] -= 1
  			if cycles[i] == 0 {
  				index := indices[i]
  				for j := i; j < n-1; j += 1 {
  					indices[j] = indices[j+1]
  				}
  				indices[n-1] = index
  				cycles[i] = n - i
  			} else {
  				j := cycles[i]
  				indices[i], indices[n-j] = indices[n-j], indices[i]

  				for k := i; k < r; k += 1 {
  					result[k] = pool[indices[k]]
  				}

  				ch <- result

  				break
  			}
  		}

  		if i < 0 {
  			ch <- nil
  		}

  	}
  }()

  return ch
}
