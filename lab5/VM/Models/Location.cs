using System.ComponentModel.DataAnnotations;

namespace VM.Models;

public class Location
{
    public int Id { get; set; }
    [StringLength(10)]
    public required string LocationId { get; set; }
    public required string Description { get; set; }
    public LocationType LocationType { get; set; }
}

public enum LocationType
{
    Forest = 0,
    Desert = 1,
    Dungeon = 2,
    River = 3,
    Ocean = 4
}
