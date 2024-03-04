using System.ComponentModel.DataAnnotations;

namespace VM.Models;

public class Messages
{
    public int Id { get; set; }
    public required Player PlayerFrom { get; set; }
    public required Player PlayerTo { get; set; }
    [StringLength(1000)]
    public string? MessageText { get; set; }
}
