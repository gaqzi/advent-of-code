package day01_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"advent-of-code/day01"
)

var (
	haystack = []int64{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	bigHaystack = []int64{
		1768, 1847, 1905, 1713, 1826, 1846, 1824, 1976, 1687, 1867, 1665, 1606,
		1946, 1886, 1858, 346, 1739, 1752, 1700, 1922, 1865, 1609, 1617, 1932,
		1346, 1213, 1933, 834, 1598, 1191, 1979, 1756, 1216, 1820, 1792, 1537,
		1341, 1390, 1709, 1458, 1808, 1885, 1679, 1977, 1869, 1614, 1938, 1622,
		1868, 1844, 1969, 1822, 1510, 1994, 1337, 1883, 1519, 1766, 1554, 1825,
		1828, 1972, 1380, 1878, 1345, 1469, 1794, 1898, 1805, 1911, 1913, 1910,
		1318, 1862, 1921, 1753, 1823, 1896, 1316, 1381, 1430, 1962, 1958, 1702,
		1923, 1993, 1789, 2002, 1788, 1970, 1955, 1887, 1870, 225, 1696, 1975,
		699, 294, 1605, 1500, 1777, 1750, 1857, 1540, 1329, 1974, 1947, 1516,
		1925, 1945, 350, 1669, 1775, 1536, 1871, 1917, 1249, 1971, 2009, 1585,
		1986, 1701, 1832, 1754, 1195, 1697, 1941, 1919, 2006, 1667, 1816, 1765,
		1631, 2003, 1861, 1000, 1791, 1786, 1843, 1939, 1951, 269, 1790, 1895,
		1355, 1833, 1466, 1998, 1806, 1881, 1234, 1856, 1619, 1727, 1874, 1877,
		195, 1783, 1797, 2010, 1764, 1863, 1852, 1841, 1892, 1562, 1650, 1942,
		1695, 1730, 1965, 1632, 1981, 1900, 1991, 1884, 1278, 1062, 1394, 1999,
		2000, 1827, 1873, 1926, 1434, 1802, 1579, 1879, 1671, 1549, 1875, 1838,
		1338, 1864, 1718, 1800, 1928, 1749, 1990, 1705,
	}
)

func TestFindSum(t *testing.T) {
	t.Run("small stack finds the right numbers", func(t *testing.T) {

		values, err := day01.FindSum(2020, haystack)

		assert.NilError(t, err, "got an error finding sum")
		assert.DeepEqual(t, []int64{1721, 299}, values)
	})

	t.Run("big stack finds the summed value", func(t *testing.T) {
		values, err := day01.FindSum(2020, bigHaystack)

		assert.NilError(t, err, "failed to find correct values")

		t.Logf("%d * %d = %d", values[0], values[1], values[0]*values[1])
	})
}

func TestFindSumSplits(t *testing.T) {
	t.Run("small stack finds the right numbers", func(t *testing.T) {
		values, err := day01.FindSumSplits(2020, 3, haystack)

		assert.NilError(t, err, "got an error finding sum")
		assert.DeepEqual(t, []int64{979, 366, 675}, values)
	})

	t.Run("big stack finds the summed value", func(t *testing.T) {
		values, err := day01.FindSumSplits(2020, 3, bigHaystack)

		assert.NilError(t, err, "got an error finding sum")
		t.Logf("%d * %d * %d = %d", values[0], values[1], values[2], values[0]*values[1]*values[2])
	})
}

var benchResult []int64

func BenchmarkFindSum(b *testing.B) {
	var values []int64

	for n:= 0; n < b.N;n++ {
		values, _ = day01.FindSum(2020, bigHaystack)
	}

	benchResult = values
}

func BenchmarkFindSumSorted(b *testing.B) {
	var values []int64

	for n:= 0; n < b.N;n++ {
		values, _ = day01.FindSumSorted(2020, bigHaystack)
	}

	benchResult = values
}

func BenchmarkFindSumMap(b *testing.B) {
	var values []int64

	for n:= 0; n < b.N;n++ {
		values, _ = day01.FindSumMap(2020, bigHaystack)
	}

	benchResult = values
}

func BenchmarkFindSumSplits(b *testing.B) {
	var values []int64

	for n:= 0; n < b.N;n++ {
		values, _ = day01.FindSumSplits(2020, 3, bigHaystack)
	}

	benchResult = values
}

func BenchmarkFindSumSplitsSorted(b *testing.B) {
	var values []int64

	for n:= 0; n < b.N;n++ {
		values, _ = day01.FindSumSplitsSorted(2020, 3, bigHaystack)
	}

	benchResult = values
}

func BenchmarkFindSumSplitsMap(b *testing.B) {
	var values []int64

	for n:= 0; n < b.N;n++ {
		values, _ = day01.FindSumSplitsMap(2020, 3, bigHaystack)
	}

	benchResult = values
}
