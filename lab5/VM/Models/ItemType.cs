using System.ComponentModel.DataAnnotations;

namespace VM.Models;

public class ItemType
{
    public int Id { get; set; }
    [StringLength(50)]
    public string? Name { get; set; }
}
