defmodule StringSeries do
  @doc """
  Given a string `s` and a positive integer `size`, return all substrings
  of that size. If `size` is greater than the length of `s`, or less than 1,
  return an empty list.
  """
  @spec slices(s :: String.t(), size :: integer) :: list(String.t())
  def slices(s, size) do
    series(s, size, [])
  end

  defp series(s, size, acc) do
    cond do
      size < 1 -> acc
      size > String.length(s) -> acc
      true -> series(String.slice(s, 1..-1), size, acc ++ [String.slice(s, 0..size-1)])
    end
  end
end

