defmodule Acronym do
  @regex ~r/[A-Z]|(?<= )[a-z]/

  @doc """
  Generate an acronym from a string.
  "This is a string" => "TIAS"
  """
  @spec abbreviate(String.t()) :: String.t()
  def abbreviate(string) do
    Regex.scan(@regex, string)
      |> Enum.join
      |> String.upcase
  end
end
