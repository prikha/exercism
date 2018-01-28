defmodule Bob do
  @sure "Sure."
  @yell "Whoa, chill out!"
  @calm_down "Calm down, I know what I'm doing!"
  @fine "Fine. Be that way!"
  @whatever "Whatever."

  @doc """
  Bob answers 'Sure.' if you ask him a question.

  He answers 'Whoa, chill out!' if you yell at him.

  He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

  He says 'Fine. Be that way!' if you address him without actually saying
  anything.

  He answers 'Whatever.' to anything else.
  """
  @spec hey(input :: String.t()) :: String.t()
  def hey(input) do
    cond do
        nothing?(input) -> @fine
        question_yell?(input) -> @calm_down
        question?(input) -> @sure
        yell?(input) -> @yell
        true -> @whatever
    end
  end

  def question?(input) do
    String.match?(input, ~r/.*\?$/)
  end

  def yell?(input) do
    String.match?(input, ~r/[0-9A-ZА-Я ]+!*$/) &&
    String.match?(input, ~r/[A-ZА-Я]/)
  end

  def question_yell?(input) do
    question?(input) && yell?(input)
  end

  def nothing?(input) do
    String.match?(input, ~r/^\s*$/)
  end
end
