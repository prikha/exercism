defmodule SecretHandshake do
  use Bitwise, only_operators: true
  @commands {"wink", "double blink", "close your eyes", "jump"}

  @doc """
  Determine the actions of a secret handshake based on the binary
  representation of the given `code`.

  If the following bits are set, include the corresponding action in your list
  of commands, in order from lowest to highest.

  1 = wink
  10 = double blink
  100 = close your eyes
  1000 = jump

  10000 = Reverse the order of the operations in the secret handshake
  """
  @spec commands(code :: integer) :: list(String.t())
  def commands(code) do
    code
      |> Integer.digits(2)
      |> Enum.reverse
      |> Enum.with_index
      |> Enum.reduce([], fn({el, idx}, acc)-> add(idx, el, acc) end)
      |> Enum.reverse
  end

  @spec add(idx :: integer, bit :: integer, accumulator :: list(String.t())) :: list(String.t())
  defp add(idx, 1, acc) when idx > 3 do
    Enum.reverse(acc)
  end

  @spec add(idx :: integer, bit :: integer, accumulator :: list(String.t())) :: list(String.t())
  defp add(_, 0, acc) do
    acc
  end

  @spec add(idx :: integer, bit :: integer, accumulator :: list(String.t())) :: list(String.t())
  defp add(idx, 1, acc) do
   [elem(@commands, idx) | acc]
  end
end

