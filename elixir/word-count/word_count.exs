defmodule Words do
  @skip ~r/[\:\,\.\!\&\@\$\%\^\&]+/
  @split ~r/[_]+/
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t) :: map
  def count(sentence) do
    sentence
      |> String.replace(@skip, "")
      |> String.replace(@split, " ")
      |> String.replace(~r/[\s]+/, " ")
      |> String.downcase
      |> String.split
      |> word_count %{}
  end

  defp word_count([], acc) do
    acc
  end

  defp word_count([word | tail], acc) do
    word_count tail, Map.put(acc, word, Map.get(acc, word, 0) + 1)
  end
end
