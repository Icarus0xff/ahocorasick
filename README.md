# Aho-Corasick
A concurrency-safe Aho-Corasick string matching algorithms implementation in golang.

## Example
```golang
    m := NewMatcher()
	dictionary := []string{"Philosopher", "Philosophe", "Philosoph", "Philosop"}
	m.Build(dictionary)
	m.Insert("Philoso")
	m.Insert("Philos")
	m.Insert("Philo")
	m.Insert("Phil")
	m.Insert("Phi")
	m.Insert("Ph")
	m.Insert("P")
	hits := m.Search("The various modes of worship, which prevailed in the Roman world, were all considered by the people, as equally true; by the Philosopher, as equally false; and by the magistrate, as equally useful.")
	fmt.Println(hits)
    
    //output:
    //map[Philos:[130] Philoso:[131] Philosop:[132] Philosoph:[133] Philosophe:[134] Philosopher:[135] P:[125] Ph:[126] Philo:[129] Phi:[127] Phil:[128]]
```
