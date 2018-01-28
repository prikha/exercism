defmodule RotationalCipher do
  @alphabet_lower for n <- ?a..?z, do: << n :: utf8 >>

  @alphabet_upper for n <- ?A..?Z, do: << n :: utf8 >>

  @alphabet_length length(@alphabet_lower)

  @doc """
  Given a plaintext and amount to shift by, return a rotated string.

  Example:
  iex> RotationalCipher.rotate("Attack at dawn", 13)
  "Nggnpx ng qnja"
  """
  @spec rotate(text :: String.t(), shift :: integer) :: String.t()
  def rotate(text, shift) do
    text
      |> String.graphemes
      |> Enum.map(fn(char)-> rotate_char(char, shift) end)
      |> Enum.join("")

  end

  defp rotate_char(lowercase, shift) when lowercase in @alphabet_lower do
    rotate_in_alphabet(@alphabet_lower, lowercase, shift)
  end

  defp rotate_char(uppercase, shift) when uppercase in @alphabet_upper do
    rotate_in_alphabet(@alphabet_upper, uppercase, shift)
  end

  defp rotate_char(non_letter, _) do
    non_letter
  end

  defp rotate_in_alphabet(alphabet, letter, shift) do
    idx = alphabet
      |> Enum.find_index(fn(x) -> x == letter end)
      |> apply_offset(shift)

    Enum.fetch!(alphabet, idx)
  end

  defp apply_offset(idx, shift) do
    rem((idx + shift), @alphabet_length)
  end
end

