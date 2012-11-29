using System;
using System.Linq;
public class Program
{
	public static void Main(String[] args)
	{
		string line;
		Console.WriteLine("Enter n t (press CTRL+Z to exit):\n Example: 10 5 => {1, 2, 3 , .... 10}, 5");
		while(true)
		{
			line = Console.ReadLine();
			if(line == null)
				break;
			int num_elems = Int32.Parse(line.Split(' ')[0]);
			int find = Int32.Parse(line.Split(' ')[1]);

			Console.WriteLine(BinarySearch(Enumerable.Range(1, num_elems).ToArray(), find));			
		}
	}

	private static int BinarySearch(int[] x, int t)
	{
		int l = 0, u = x.Length - 1, m;
		while (l <= u)
		{
			m = (l + u) / 2;
			if(x[m] > t)
				u = m - 1;
			else if(x[m] < t)
				l = m + 1;
			else if(x[m] == t)
				return m;
		}
		return -1;
	}

	private static bool isSorted(int[] x)
	{
		for(int i=0; i < x.Length - 1; i ++)
			if(x[i] > x[i+1])
				return false;
		return true;
	}
}
