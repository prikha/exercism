defmodule PigLatin do
  @vowels ["a", "e", "i", "o", "u", "yt", "xr"]

  @doc """
  Given a `phrase`, translate it a word at a time to Pig Latin.

  Words beginning with consonants should have the consonant moved to the end of
  the word, followed by "ay".

  Words beginning with vowels (aeiou) should have "ay" added to the end of the
  word.

  Some groups of letters are treated like consonants, including "ch", "qu",
  "squ", "th", "thr", and "sch".

  Some groups are treated like vowels, including "yt" and "xr".
  """
  @spec translate(phrase :: String.t()) :: String.t()
  def translate(phrase) do
    String.split(phrase)
      |> Enum.map(&translate_word(&1))
      |> Enum.join(" ")
  end

  defp translate_word(word) do
    cond do
      String.starts_with?(word, @vowels) ->
        word <> "ay"

      String.starts_with?(word, "qu") ->
        rotate_sound("qu", word)
        |> translate_word()

      true ->
        char = String.first(word)
        rotate_sound(char, word)
        |> translate_word()
    end
  end

  defp rotate_sound(prefix, word) do
    String.replace_prefix(word, prefix, "") <> prefix
  end
end