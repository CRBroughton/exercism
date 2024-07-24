import gleam/int
import gleam/option.{type Option, None, Some}

pub type Player {
  Player(name: Option(String), level: Int, health: Int, mana: Option(Int))
}

pub fn introduce(player: Player) -> String {
  case player.name {
    Some(name) -> name
    None -> "Mighty Magician"
  }
}

pub fn revive(player: Player) -> Option(Player) {
  case player.health, player.level {
    0, level if level > 9 ->
      Some(Player(..player, health: 100, mana: Some(100)))
    0, _ -> Some(Player(..player, health: 100, mana: player.mana))
    _, _ -> None
  }
}

pub fn cast_spell(player: Player, cost: Int) -> #(Player, Int) {
  case player.mana {
    Some(mana) if mana >= cost -> #(
      Player(..player, mana: Some(mana - cost)),
      2 * cost,
    )
    None -> #(Player(..player, health: int.max(0, player.health - cost)), 0)
    _ -> #(player, 0)
  }
}
