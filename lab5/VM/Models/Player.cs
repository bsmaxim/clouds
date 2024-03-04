using System.ComponentModel.DataAnnotations;

namespace VM.Models;

public class Player
{
    public int Id { get; set; }
    [StringLength(30)]
    public required string Name { get; set; }
    public PlayerClass PlayerClass { get; set; }
    [EmailAddress]
    public string? Email { get; set; }
    public int Level { get; set; }
    public required string Position { get; set; }
}

public enum PlayerClass
{
    Knight = 0,
    Wizard = 1,
    Thief = 2,
    Paladin = 3
}
