package powerbill

import "fmt"

// Consumption рассчитывает объем потребленной электроэнергии за период.
// Принимает предыдущие (prev) и текущие (curr) показания счетчика в кВт·ч.
// Возвращает ошибку, если текущие показания меньше предыдущих (curr < prev)
// или если хотя бы одно из значений отрицательное.
func Consumption(prev, curr float64) (float64, error) {
	if prev < 0.0 {
		return 0.0, fmt.Errorf("входные данные не могут быть отрицательными (prev: %.2f)", prev)
	}

	if curr < 0.0 {
		return 0.0, fmt.Errorf("входные данные не могут быть отрицательными (curr: %.2f)", curr)
	}

	if curr < prev {
		return 0.0, fmt.Errorf("текущий показатель должен быть >= предыдущего (prev: %.2f, curr: %.2f)", prev, curr)
	}

	return curr - prev, nil
}

// EnergyCost рассчитывает базовую стоимость электроэнергии.
// Принимает объем энергии (kwh) и тариф за 1 кВт·ч (tariff).
// Возвращает ошибку, если kwh или tariff имеют отрицательное значение.
func EnergyCost(kwh, tariff float64) (float64, error) {
	if kwh < 0.0 {
		return 0.0, fmt.Errorf("входные данные не могут быть отрицательными (kwh: %.2f)", kwh)
	}

	if tariff < 0.0 {
		return 0.0, fmt.Errorf("входные данные не могут быть отрицательными (tariff: %.2f)", tariff)
	}

	return kwh * tariff, nil
}

// ApplyDiscount изменяет стоимость электроэнергии (cost) по указателю,
// применяя скидку в процентах (percent).
// Возвращает ошибку, если cost равен nil или если percent не входит в диапазон [0, 100].
func ApplyDiscount(cost *float64, percent float64) error {
	if cost == nil {
		return fmt.Errorf("цена должна быть указана")
	}

	if percent < 0 || percent > 100 {
		return fmt.Errorf("несоответствие законам физики (percent: %.2f)", percent)
	}

	*cost *= (1 - percent/100)
	return nil
}

// FormatEnergyReport формирует строковый отчет о начислениях за электроэнергию.
// Принимает имя владельца (owner), потребление (kwh) и итоговую стоимость (cost).
// Возвращает ошибку, если owner пуст, а также если kwh или cost меньше нуля.
func FormatEnergyReport(owner string, kwh, cost float64) (string, error) {
	if owner == "" {
		return "", fmt.Errorf("налоговая работает только с людьми, укажите имя")
	}

	if kwh < 0.0 {
		return "", fmt.Errorf("входные данные не могут быть отрицательными (kwh: %.2f)", kwh)
	}

	if cost < 0.0 {
		return "", fmt.Errorf("входные данные не могут быть отрицательными (cost: %.2f)", cost)
	}

	report := fmt.Sprintf(
		"========================================\n"+
			"           ЭНЕРГОУЧЕТ // ОТЧЕТ           \n"+
			"========================================\n"+
			" АБОНЕНТ:     %s\n"+
			" РАСХОД:      %.2f кВт·ч\n"+
			" -------------------------------------- \n"+
			" ИТОГО:       %.2f руб.\n"+
			"========================================",
		owner, kwh, cost,
	)

	return report, nil
}
