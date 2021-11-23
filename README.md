# Custom text condition for filters

## Install

    go get -v github.com/blins/condition

## Use

    func init() {
        condition.RegisterConditionFabric("prefix", condition.StringCheckerFabric(strings.HasPrefix))
        condition.RegisterConditionFabric("suffix", condition.StringCheckerFabric(strings.HasSuffix))

    }

    func main() {
        cmdLine := strings.Fields("prefix start and suffix stop")
        cond, _ := condition.ParseArgs(cmdLine)

        cond.Check("superstart") // false
        cond.Check("startone") // false
        cond.Check("start and other text stop") // true
    }

Example with keyword in folder `example`

    